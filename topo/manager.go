package topo

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	topologyclientv1 "github.com/networkop/meshnet-cni/api/clientset/v1beta1"
	topologyv1 "github.com/networkop/meshnet-cni/api/types/v1beta1"
	ktpb "github.com/openconfig/kne/proto/topo"
	log "github.com/sirupsen/logrus"
	"github.com/v3rgilius/bgpemu/helper"
	tpb "github.com/v3rgilius/bgpemu/proto/bgptopo"
	node "github.com/v3rgilius/bgpemu/topo/node"
	_ "github.com/v3rgilius/bgpemu/topo/node/gobgp"
	_ "github.com/v3rgilius/bgpemu/topo/node/host"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"

	"google.golang.org/protobuf/encoding/prototext"
)

type UpdateManager struct {
	topo       *tpb.Topology
	ktopo      *ktpb.Topology
	nodes      map[string]node.Node
	kubecfg    string
	kClient    kubernetes.Interface
	tClient    topologyclientv1.Interface
	rCfg       *rest.Config
	dInterface dynamic.NamespaceableResourceInterface
	basePath   string
}

var gvr = schema.GroupVersionResource{
	Group:    topologyv1.GroupName,
	Version:  topologyv1.GroupVersion,
	Resource: "topologies",
}

func GVR() schema.GroupVersionResource {
	return gvr
}

var (
	groupVersion = &schema.GroupVersion{
		Group:   topologyv1.GroupName,
		Version: topologyv1.GroupVersion,
	}
)

func GV() *schema.GroupVersion {
	return groupVersion
}

func New(topo *tpb.Topology, ktopo *ktpb.Topology, startUid int64) (*UpdateManager, error) {
	if ktopo == nil {
		return nil, fmt.Errorf("topology cannot be nil")
	}
	m := &UpdateManager{
		topo:    topo,
		ktopo:   ktopo,
		nodes:   map[string]node.Node{},
		kubecfg: helper.DefaultKubeCfg(),
	}
	if m.rCfg == nil {
		log.Infof("Trying in-cluster configuration")
		rCfg, err := rest.InClusterConfig()
		if err != nil {
			log.Infof("Falling back to kubeconfig: %q", m.kubecfg)
			rCfg, err = clientcmd.BuildConfigFromFlags("", m.kubecfg)
			if err != nil {
				return nil, err
			}
		}
		m.rCfg = rCfg
	}
	if m.kClient == nil {
		kClient, err := kubernetes.NewForConfig(m.rCfg)
		if err != nil {
			return nil, err
		}
		m.kClient = kClient
	}
	if m.tClient == nil {
		tClient, err := topologyclientv1.NewForConfig(m.rCfg)
		if err != nil {
			return nil, err
		}
		m.tClient = tClient
		dClient, err := dynamic.NewForConfig(m.rCfg)
		if err != nil {
			return nil, err
		}
		dInterface := dClient.Resource(gvr)
		m.dInterface = dInterface
	}
	if err := m.load(startUid); err != nil {
		return nil, fmt.Errorf("failed to load topology: %w", err)
	}
	log.Infof("Created manager for topology:\n%v", prototext.Format(m.ktopo))
	return m, nil
}

// Create creates the topology in the cluster.
func (m *UpdateManager) Create(ctx context.Context, timeout time.Duration) error {
	log.Infof("Topology:\n%v", prototext.Format(m.ktopo))
	if err := m.push(ctx); err != nil {
		return err
	}
	if err := m.checkNodeStatus(ctx, timeout); err != nil {
		return err
	}
	log.Infof("Topology %q created", m.ktopo.GetName())
	return nil
}

func Update(t *tpb.Topology) error {
	kt, err := KneTopo(t)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	bt, err := BaseTopo(t.UpdateTopo)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	kbt, _ := KneTopo(bt)

	tnodes := make(map[string]*ktpb.Node, 64)
	btnodes := make(map[string]*ktpb.Node, 64)
	newLinks := make([]*ktpb.Link, 0, 64)
	existedMTs := make(map[string][]*topologyv1.Link, 16)
	newMTs := make(map[string][]*topologyv1.Link, 16)
	curUid := len(bt.Links) + len(t.Links)
	for _, n := range kbt.Nodes {
		btnodes[n.Name] = n
	}
	for _, n := range kt.Nodes {
		tnodes[n.Name] = n
		if ok := btnodes[n.Name]; ok != nil {
			return fmt.Errorf("Can't update existed pod: %s", n.Name)
		}
	}
	for _, l := range kt.Links {
		flag := false
		flagA := false
		flagZ := false
		if ok := btnodes[l.ANode]; ok != nil {
			flag = true
			addMeshTopoNode(existedMTs, true, curUid, l)
		} else {
			flagA = true
		}
		if ok := btnodes[l.ZNode]; ok != nil {
			flag = true
			addMeshTopoNode(existedMTs, false, curUid, l)
		} else {
			flagZ = true
		}
		if flagA && flag {
			addMeshTopoNode(newMTs, true, curUid, l)
		}
		if flagZ && flag {
			addMeshTopoNode(newMTs, false, curUid, l)
		}
		if !flag {
			newLinks = append(newLinks, l)
		} else {
			curUid++
		}
	}
	kt.Links = newLinks
	tm, err := New(t, kt, int64(len(bt.Links)))
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	err = tm.updateMeshnetTopologies(context.Background(), existedMTs, newMTs)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	err = tm.Create(context.Background(), 0)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	// err = UpdatePods(t, tm)
	// if err != nil {
	// 	return fmt.Errorf("%w", err)
	// }
	return nil
}

func addMeshTopoNode(mts map[string][]*topologyv1.Link, isA bool, uid int, link *ktpb.Link) {
	var nodeName, aInt, zNode, zInt string
	if isA {
		nodeName = link.ANode
		aInt = link.AInt
		zNode = link.ZNode
		zInt = link.ZInt
	} else {
		nodeName = link.ZNode
		aInt = link.ZInt
		zNode = link.ANode
		zInt = link.AInt
	}
	if mt, ok := mts[nodeName]; ok {
		mts[nodeName] = []*topologyv1.Link{
			{
				LocalIntf: aInt,
				PeerPod:   zNode,
				PeerIntf:  zInt,
				UID:       uid,
			},
		}
	} else {
		mts[nodeName] = append(mt, &topologyv1.Link{
			LocalIntf: aInt,
			PeerPod:   zNode,
			PeerIntf:  zInt,
			UID:       uid,
		})
	}
}

// Delete deletes the topology from the cluster.
func (m *UpdateManager) Delete(ctx context.Context) error {
	log.Infof("Topology:\n%v", prototext.Format(m.ktopo))
	if _, err := m.kClient.CoreV1().Namespaces().Get(ctx, m.ktopo.Name, metav1.GetOptions{}); err != nil {
		return fmt.Errorf("topology %q does not exist in cluster", m.ktopo.Name)
	}

	if err := m.deleteMeshnetTopologies(ctx); err != nil {
		log.Errorf("%s", err)
		// return err
	}

	// Delete topology nodes
	for _, n := range m.nodes {
		// Delete Service for node
		if err := n.Delete(ctx); err != nil {
			log.Warningf("Error deleting node %q: %v", n.Name(), err)
		}
	}

	// Delete namespace
	prop := metav1.DeletePropagationForeground
	log.Infof("Topology %q deleted", m.ktopo.GetName())
	return m.kClient.CoreV1().Namespaces().Delete(ctx, m.ktopo.Name, metav1.DeleteOptions{PropagationPolicy: &prop})
}

func (m *UpdateManager) GenerateSelfSigned(ctx context.Context, nodeName string) error {
	n, ok := m.nodes[nodeName]
	if !ok {
		return fmt.Errorf("node %q not found", nodeName)
	}
	if n.GetProto().GetConfig().GetCert() == nil {
		log.Infof("No cert info for %q, skipping cert generation", nodeName)
		return nil
	}
	c, ok := n.(node.Certer)
	if !ok {
		return status.Errorf(codes.Unimplemented, "node %q does not implement Certer interface", nodeName)
	}
	return c.GenerateSelfSigned(ctx)
}

func (m *UpdateManager) checkNodeStatus(ctx context.Context, timeout time.Duration) error {
	foundAll := false
	processed := make(map[string]bool)

	// Check until end state or timeout sec expired
	start := time.Now()
	for (timeout == 0 || time.Since(start) < timeout) && !foundAll {
		foundAll = true
		for name, n := range m.nodes {
			if _, ok := processed[name]; ok {
				continue
			}

			phase, err := n.Status(ctx)
			if err != nil || phase == node.StatusFailed {
				return fmt.Errorf("Node %q: Status %s Reason %v", name, phase, err)
			}
			if phase == node.StatusRunning {
				log.Infof("Node %q: Status %s", name, phase)
				processed[name] = true
			} else {
				foundAll = false
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
	if !foundAll {
		log.Warningf("Failed to determine status of some node resources in %d sec", timeout)
	}
	for name, n := range m.nodes {
		tasks := n.GetOpt().Tasks
		for _, task := range tasks {
			_ = m.Exec(ctx, task.Cmds, name, task.Container, nil, os.Stdout, os.Stderr)
		}
	}
	return nil
}

// push deploys the topology to the cluster.
func (m *UpdateManager) push(ctx context.Context) error {
	if _, err := m.kClient.CoreV1().Namespaces().Get(ctx, m.ktopo.Name, metav1.GetOptions{}); err != nil {
		log.Infof("Creating namespace for topology: %q", m.ktopo.Name)
		ns := &corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: m.ktopo.Name,
			},
		}
		sNs, err := m.kClient.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
		if err != nil {
			return err
		}
		log.Infof("Server Namespace: %+v", sNs)
	}

	if err := m.createMeshnetTopologies(ctx); err != nil {
		return err
	}

	log.Infof("Creating Node Pods")
	for k, n := range m.nodes {
		if err := n.Create(ctx); err != nil {
			return err
		}
		log.Infof("Node %q resource created", k)
	}
	for _, n := range m.nodes {
		err := m.GenerateSelfSigned(ctx, n.Name())
		switch {
		default:
			return fmt.Errorf("failed to generate cert for node %s: %w", n.Name(), err)
		case err == nil, status.Code(err) == codes.Unimplemented:
		}
	}
	return nil
}

// createMeshnetTopologies creates meshnet resources for all available nodes.
func (m *UpdateManager) createMeshnetTopologies(ctx context.Context) error {
	log.Infof("Getting topology specs for namespace %s", m.ktopo.Name)
	topologies, err := m.topologySpecs(ctx)
	if err != nil {
		return fmt.Errorf("could not get meshnet topologies: %v", err)
	}
	log.Infof("Got topology specs for namespace %s: %+v", m.ktopo.Name, topologies)
	for _, t := range topologies {
		log.Infof("Creating topology for meshnet node %s", t.ObjectMeta.Name)
		sT, err := m.tClient.Topology(m.ktopo.Name).Create(ctx, t, metav1.CreateOptions{})
		if err != nil {
			return fmt.Errorf("could not create topology for meshnet node %s: %v", t.ObjectMeta.Name, err)
		}
		log.Infof("Meshnet Node:\n%+v\n", sT)
	}
	return nil
}

// createMeshnetTopologies creates meshnet resources for all available nodes.
// update existing meshnet resources AND modify nodes (link's UID) in manager
func (m *UpdateManager) updateMeshnetTopologies(ctx context.Context, updateMTs map[string][]*topologyv1.Link, newMTs map[string][]*topologyv1.Link) error {
	for name, links := range updateMTs {
		mt, err := m.tClient.Topology(m.ktopo.Name).Get(ctx, name, metav1.GetOptions{})
		if err != nil {
			return err
		}
		for _, link := range links {
			mt.Spec.Links = append(mt.Spec.Links, *link)
		}
		obj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(mt)
		if err != nil {
			return err
		}
		_, err = m.dInterface.Update(ctx, &unstructured.Unstructured{Object: obj}, metav1.UpdateOptions{})
		if err != nil {
			return err
		}
	}

	for name, links := range newMTs {
		n := m.nodes[name].GetProto()
		for _, l := range links {
			n.Interfaces[l.LocalIntf] = &ktpb.Interface{
				PeerName:    l.PeerPod,
				PeerIntName: l.PeerIntf,
				Uid:         int64(l.UID),
				IntName:     l.LocalIntf,
			}
		}
	}
	return nil
}

// deleteMeshnetTopologies deletes meshnet resources for all available nodes.
func (m *UpdateManager) deleteMeshnetTopologies(ctx context.Context) error {
	nodes, err := m.topologyResources(ctx)
	if err == nil {
		for _, n := range nodes {
			if err := m.tClient.Topology(m.ktopo.Name).Delete(ctx, n.ObjectMeta.Name, metav1.DeleteOptions{}); err != nil {
				log.Warningf("Error meshnet node %q: %v", n.ObjectMeta.Name, err)
			}
		}
	} else {
		// no need to return warning as deleting meshnet namespace shall delete the resources too
		log.Warningf("Error getting meshnet nodes: %v", err)
	}

	return nil
}

// load populates the internal fields of the topology proto.
func (m *UpdateManager) load(startUid int64) error {
	nMap := map[string]*ktpb.Node{}
	for _, n := range m.ktopo.Nodes {
		if len(n.Interfaces) == 0 {
			n.Interfaces = map[string]*ktpb.Interface{}
		}
		for k := range n.Interfaces {
			if n.Interfaces[k].IntName == "" {
				n.Interfaces[k].IntName = k
			}
		}
		nMap[n.Name] = n
	}
	nConfig := make(map[string]*tpb.Config, 16)
	for _, n := range m.topo.Nodes {
		ipcmds := make([]string, 0, 8)
		for eth, ip := range n.IpAddr {
			ipcmds = append(ipcmds, fmt.Sprintf("ip addr add %s dev %s", ip, eth))
		}
		n.Config.Tasks = append(n.Config.Tasks, &tpb.Task{
			Container: n.Name,
			Cmds:      ipcmds,
		})
		nConfig[n.Name] = n.Config
	}
	uid := startUid
	for _, l := range m.ktopo.Links {
		log.Infof("Adding Link: %s:%s %s:%s", l.ANode, l.AInt, l.ZNode, l.ZInt)
		aNode, ok := nMap[l.ANode]
		if !ok {
			return fmt.Errorf("invalid topology: missing node %q", l.ANode)
		}
		aInt, ok := aNode.Interfaces[l.AInt]
		if !ok {
			aInt = &ktpb.Interface{
				IntName: l.AInt,
			}
			aNode.Interfaces[l.AInt] = aInt
		}
		zNode, ok := nMap[l.ZNode]
		if !ok {
			return fmt.Errorf("invalid topology: missing node %q", l.ZNode)
		}
		zInt, ok := zNode.Interfaces[l.ZInt]
		if !ok {
			zInt = &ktpb.Interface{
				IntName: l.ZInt,
			}
			zNode.Interfaces[l.ZInt] = zInt
		}
		if aInt.PeerName != "" {
			return fmt.Errorf("interface %s:%s already connected", l.ANode, l.AInt)
		}
		if zInt.PeerName != "" {
			return fmt.Errorf("interface %s:%s already connected", l.ZNode, l.ZInt)
		}
		aInt.PeerName = l.ZNode
		aInt.PeerIntName = l.ZInt
		aInt.Uid = uid
		zInt.PeerName = l.ANode
		zInt.PeerIntName = l.AInt
		zInt.Uid = uid
		uid++
	}
	for k, n := range nMap {
		log.Infof("Adding Node: %s:%s", n.Name, n.Vendor)
		nn, err := node.New(m.ktopo.Name, n, m.kClient, m.rCfg, m.basePath, m.kubecfg, nConfig[k])
		if err != nil {
			return fmt.Errorf("failed to load topology: %w", err)
		}
		m.nodes[k] = nn
	}
	return nil
}

// topologySpecs provides a custom implementation for constructing meshnet resource specs
// (before meshnet topology creation) for all configured nodes.
func (m *UpdateManager) topologySpecs(ctx context.Context) ([]*topologyv1.Topology, error) {
	nodeSpecs := map[string][]*topologyv1.Topology{}
	topos := []*topologyv1.Topology{}

	// get topology specs from all nodes
	for _, n := range m.nodes {
		log.Infof("Getting topology specs for node %s", n.Name())
		specs, err := n.TopologySpecs(ctx)
		if err != nil {
			return nil, fmt.Errorf("could not fetch topology specs for node %s: %v", n.Name(), err)
		}

		log.Infof("Topology specs for node %s: %+v", n.Name(), specs)
		nodeSpecs[n.Name()] = specs
	}

	// replace node name with pod name, for peer pod attribute in each link
	for nodeName, specs := range nodeSpecs {
		for _, spec := range specs {
			for l := range spec.Spec.Links {
				link := &spec.Spec.Links[l]
				peerSpecs, ok := nodeSpecs[link.PeerPod]
				if !ok {
					return nil, fmt.Errorf("specs do not exist for node %s", link.PeerPod)
				}

				if err := setLinkPeer(nodeName, spec.ObjectMeta.Name, link, peerSpecs); err != nil {
					return nil, err
				}
			}
			topos = append(topos, spec)
		}
	}

	return topos, nil
}

// topologyResources gets the topology CRDs for the cluster.
func (m *UpdateManager) topologyResources(ctx context.Context) ([]*topologyv1.Topology, error) {
	topology, err := m.tClient.Topology(m.ktopo.Name).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get topology CRDs: %v", err)
	}

	items := make([]*topologyv1.Topology, len(topology.Items))
	for i := range items {
		items[i] = &topology.Items[i]
	}

	return items, nil
}

// setLinkPeer finds the peer pod name and peer interface name for a given interface.
func setLinkPeer(nodeName string, podName string, link *topologyv1.Link, peerSpecs []*topologyv1.Topology) error {
	for _, peerSpec := range peerSpecs {
		for _, peerLink := range peerSpec.Spec.Links {
			// make sure self ifc and peer ifc belong to same link (and hence UID) but are not the same interfaces
			if peerLink.UID == link.UID && !(nodeName == link.PeerPod && peerLink.LocalIntf == link.LocalIntf) {
				link.PeerPod = peerSpec.ObjectMeta.Name
				link.PeerIntf = peerLink.LocalIntf
				return nil
			}
		}
	}
	return fmt.Errorf("could not find peer for node %s pod %s link UID %d", nodeName, podName, link.UID)
}

func (m *UpdateManager) Exec(ctx context.Context, cmds []string, podname string, name string, stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	for _, command := range cmds {
		cmd := []string{
			"/bin/sh",
			"-c",
			command,
		}
		req := m.kClient.CoreV1().RESTClient().Post().
			Resource("pods").
			Name(podname).
			Namespace(m.ktopo.Name).
			SubResource("exec")
		req.VersionedParams(&corev1.PodExecOptions{
			Stdin:     false,
			Stdout:    true,
			Stderr:    true,
			Container: name,
			Command:   cmd,
			TTY:       false,
		}, scheme.ParameterCodec)
		log.Infof("Executing extra commands on container %s: %s", name, command)
		exec, err := remotecommand.NewSPDYExecutor(m.rCfg, "POST", req.URL())
		if err != nil {
			log.Errorf("error in creating executor for extra commands of container %s : %s", name, err.Error())
			return err
		}
		err = exec.StreamWithContext(context.Background(), remotecommand.StreamOptions{
			Stdin:  stdin,
			Stdout: stdout,
			Stderr: stderr,
			Tty:    false,
		})
		if err != nil {
			log.Errorf("error in executing extra commands of node %s : %s", name, err.Error())
			return err
		}
	}
	return nil
}
