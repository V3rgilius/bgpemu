package topo

import (
	"context"
	"fmt"

	ktpb "github.com/openconfig/kne/proto/topo"
	tpb "github.com/v3rgilius/bgpemu/proto/bgptopo"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/prototext"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
	"os"
	"sigs.k8s.io/yaml"
	"strings"
)

var protojsonUnmarshaller = protojson.UnmarshalOptions{
	AllowPartial:   true,
	DiscardUnknown: false,
}

// Load loads a Topology from path and parse all subtopos.
func Load(path string) (*tpb.Topology, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	t := &tpb.Topology{}
	switch {
	case strings.HasSuffix(path, ".yaml"), strings.HasSuffix(path, ".yml"):
		jsonBytes, err := yaml.YAMLToJSON(b)
		if err != nil {
			return nil, fmt.Errorf("could not parse yaml: %v", err)
		}
		if err := protojsonUnmarshaller.Unmarshal(jsonBytes, t); err != nil {
			return nil, fmt.Errorf("could not parse json: %v", err)
		}
	default:
		if err := prototext.Unmarshal(b, t); err != nil {
			return nil, err
		}
	}
	fullTopo, err := parseSubTopo(t)
	if err != nil {
		return nil, err
	}
	return fullTopo, nil
}

// parse subtopo to full nodes and links in topo,but might with export interfaces
func parseSubTopo(t *tpb.Topology) (*tpb.Topology, error) {
	subTopos := make([]*tpb.Node, 0, 16)
	subToposExports := make(map[string]map[string]*tpb.InternalInterface, 16)
	fullTopo := &tpb.Topology{
		Name:       t.Name,
		Nodes:      make([]*tpb.Node, 0, 16),
		Links:      make([]*tpb.Link, 0, 16),
		ExportInts: make(map[string]*tpb.InternalInterface, 16),
		UpdateTopo: t.UpdateTopo,
	}
	for _, node := range t.GetNodes() {
		if node.Type == tpb.Type_SUBTOPO {
			subt, err := Load(node.GetPath())
			if err != nil {
				return nil, err
			}
			fullTopo.Nodes = append(fullTopo.Nodes, subt.Nodes...)
			fullTopo.Links = append(fullTopo.Links, subt.Links...)
			subToposExports[node.Name] = subt.GetExportInts()
		} else {
			fullTopo.Nodes = append(fullTopo.Nodes, node)
		}
	}
	if len(subTopos) == 0 {
		return t, nil
	}

	for _, link := range t.GetLinks() {
		if internalInterfaces := subToposExports[link.ANode]; internalInterfaces != nil {
			link.ANode = internalInterfaces[link.AInt].Node
			link.AInt = internalInterfaces[link.AInt].NodeInt
		}
		if internalInterfaces := subToposExports[link.ZNode]; internalInterfaces != nil {
			link.ZNode = internalInterfaces[link.ZInt].Node
			link.ZInt = internalInterfaces[link.ZInt].NodeInt
		}
		fullTopo.Links = append(fullTopo.Links, link)
	}

	for intName, internalInterface := range t.ExportInts {
		if export := subToposExports[internalInterface.Node]; export != nil { //Internal node is a subtopo node
			newInternalInterface := &tpb.InternalInterface{
				Node:    export[internalInterface.NodeInt].Node,
				NodeInt: export[internalInterface.NodeInt].NodeInt,
			}
			fullTopo.ExportInts[intName] = newInternalInterface
		} else {
			fullTopo.ExportInts[intName] = internalInterface
		}
	}
	return fullTopo, nil
}

func KneTopo(t *tpb.Topology) (*ktpb.Topology, error) {
	kt := &ktpb.Topology{
		Name: t.Name,
	}
	nodes := make([]*ktpb.Node, 0, 16)
	nodeVendor := ktpb.Vendor_UNKNOWN
	bgpImg := "midnightreiter/gobgp:v3.11.0"
	hostImg := "alpine:latest"
	for _, node := range t.GetNodes() {
		kConfig := &ktpb.Config{}
		switch node.Type {
		case tpb.Type_BGPNODE:
			nodeVendor = ktpb.Vendor_GOBGP
			if node.Config.GetImage() == "" {
				kConfig.Image = bgpImg
			}
			kConfig.Command = []string{"/bin/sleep", "2000000"}
		case tpb.Type_HOST:
			nodeVendor = ktpb.Vendor_HOST
			if node.Config.GetImage() == "" {
				kConfig.Image = hostImg
			}
		case tpb.Type_SUBTOPO:
			return nil, fmt.Errorf("Subtopo!")
		}

		kneNode := &ktpb.Node{
			Name:   node.Name,
			Vendor: nodeVendor,
			Config: kConfig,
		}
		if services := node.GetServices(); services != nil {
			ksrvs := KneServices(services)
			kneNode.Services = ksrvs
		}
		nodes = append(nodes, kneNode)
	}
	kt.Nodes = nodes
	kt.Links = KneLinks(t.Links)
	return kt, nil
}

func BaseTopo(path string) (*tpb.Topology, error) {
	if path == "" {
		return nil, fmt.Errorf("UpdatePath can't be empty! ")
	}
	t, err := Load(path)
	if err != nil {
		return nil, err
	}
	if t.GetUpdateTopo() != "" {
		bt, err := BaseTopo(t.UpdateTopo)
		if err != nil {
			return nil, err
		}
		mergeTopo(bt, t)
	}
	return t, nil
}

func mergeTopo(bt *tpb.Topology, t *tpb.Topology) {
	t.Nodes = append(t.Nodes, bt.Nodes...)
	t.Links = append(t.Links, bt.Links...)
}

func UpdatePods(t *tpb.Topology, m *UpdateManager) error {
	for _, n := range t.Nodes {
		tasks := n.Config.Tasks
		name := n.Name
		for _, task := range tasks {
			err := m.Exec(context.Background(), task.Cmds, name, task.Container, nil, os.Stdout, os.Stderr)
			if err != nil {
				return err
			}
		}
		if n.Type == tpb.Type_BGPNODE {
			err := UpdateBgpPodSpec(n, m)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func UpdateBgpPodSpec(n *tpb.Node, m *UpdateManager) error {
	bgpCtnName := fmt.Sprintf("%s-frr", n.Name)
	bgpCtnImage := "frrouting/frr:v8.1.0"
	pod, err := m.kClient.CoreV1().Pods(m.ktopo.Name).Get(context.TODO(), n.Name, metav1.GetOptions{})
	if err != nil {
		return err
	}
	// 创建新容器
	newContainer := corev1.Container{
		Name:            bgpCtnName,
		Image:           bgpCtnImage,
		Command:         []string{},
		Args:            []string{},
		ImagePullPolicy: "IfNotPresent",
		SecurityContext: &corev1.SecurityContext{
			Privileged: pointer.Bool(true),
		},
	}

	// 将新容器添加到Pod的spec.containers中
	pod.Spec.Containers = append(pod.Spec.Containers, newContainer)

	shareVolumes := n.Config.ShareVolumes
	for _, sv := range shareVolumes {
		pod.Spec.Volumes = append(pod.Spec.Volumes, corev1.Volume{
			Name: fmt.Sprintf("volume-%s", sv),
			VolumeSource: corev1.VolumeSource{
				EmptyDir: &corev1.EmptyDirVolumeSource{},
			},
		})
	}
	for i, c := range pod.Spec.Containers {
		// pod.Spec.Containers[i].VolumeMounts = append(c.VolumeMounts, corev1.VolumeMount{
		// 	Name:      "zebra-volume",
		// 	MountPath: "/var/run/frr",
		// })
		if configs, ok := n.Config.ContainerVolumes[c.Name]; ok {
			for j, sv := range configs.Volumes {
				pod.Spec.Containers[i].VolumeMounts = append(c.VolumeMounts, corev1.VolumeMount{
					Name:      fmt.Sprintf("volume-%s", sv),
					MountPath: configs.Paths[j],
				})
			}
		}
	}
	// 更新Pod
	_, err = m.kClient.CoreV1().Pods(m.ktopo.Name).Update(context.TODO(), pod, metav1.UpdateOptions{})
	if err != nil {
		return err
	}

	fmt.Printf("Container %s has been added to Pod %s. \n", bgpCtnName, n.Name)
	return nil
}

func KneServices(srvs map[uint32]*tpb.Service) map[uint32]*ktpb.Service {
	ksrvs := make(map[uint32]*ktpb.Service, 4)
	for port, srv := range srvs {
		ksrvs[port] = &ktpb.Service{
			Name:   srv.Name,
			Inside: srv.Inside,
		}
	}
	return ksrvs
}

func KneLinks(links []*tpb.Link) []*ktpb.Link {
	klinks := make([]*ktpb.Link, 0, 16)
	for _, link := range links {
		klinks = append(klinks, &ktpb.Link{
			ANode: link.ANode,
			AInt:  link.AInt,
			ZNode: link.ZNode,
			ZInt:  link.ZInt,
		})
	}
	return klinks
}
