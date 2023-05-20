package topo

import (
	"fmt"
	"os"
	"strings"

	ktpb "github.com/openconfig/kne/proto/topo"
	tpb "github.com/p3rdy/bgpemu/proto/bgptopo"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/prototext"
	"sigs.k8s.io/yaml"
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
		Links:      make([]*ktpb.Link, 0, 16),
		ExportInts: make(map[string]*tpb.InternalInterface, 16),
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
func LoadToKneTopo(path string) (*ktpb.Topology, error) {
	t, err := Load(path)
	if err != nil {
		return nil, err
	}
	kt := &ktpb.Topology{
		Name: t.Name,
	}
	nodes := make([]*ktpb.Node, 0, 16)
	for _, node := range t.GetNodes() {
		nodeVendor := ktpb.Vendor_UNKNOWN
		switch node.Type {
		case tpb.Type_BGPNODE:
			nodeVendor = ktpb.Vendor_GOBGP
		case tpb.Type_HOST:
			nodeVendor = ktpb.Vendor_HOST
		case tpb.Type_SUBTOPO:
			nodeVendor = ktpb.Vendor_HOST // TODO: Parse SUBTOPO

		}

		kneNode := &ktpb.Node{
			Name:   node.Name,
			Vendor: nodeVendor,
		}
		kneConfig := &ktpb.Config{
			Tasks: []*ktpb.Task{{
				Container: node.Name,
				Cmds:      []string{},
			}},
		}
		if ipAddr := node.GetIpAddr(); ipAddr != nil {
			for eth, addr := range ipAddr {
				kneConfig.Tasks[0].Cmds = append(kneConfig.Tasks[0].Cmds, fmt.Sprintf("ip addr add %s dev %s", addr, eth))
			}
		}
		if config := node.GetConfig(); config != nil {
			kneConfig.ShareVolumes = config.GetShareVolumes()
			kneConfig.ContainerVolumes = config.GetContainerVolumes()
			kneConfig.ExtraImages = config.GetExtraImages()
			kneConfig.Tasks = append(kneConfig.Tasks, config.GetTasks()...)
		}
		if node.GetServices() != nil {
			kneNode.Services = node.GetServices()
		}
		nodes = append(nodes, kneNode)
	}
	kt.Nodes = nodes
	kt.Links = t.Links
	return kt, nil
}
