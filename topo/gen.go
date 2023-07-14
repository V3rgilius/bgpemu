package topo

import (
	"fmt"
	astopo "github.com/v3rgilius/bgpemu/proto/astopo"
	tpb "github.com/v3rgilius/bgpemu/proto/bgptopo"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/prototext"
	"os"
	"sigs.k8s.io/yaml"
	"strconv"
	"strings"
)

func GenerateFromAS(path string, outpath string) (*tpb.Topology, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	t := &astopo.Topology{}
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
	topo, err := genTopo(t)
	if err != nil {
		return nil, err
	}
	jsonbytes, err := protojson.Marshal(topo)
	if err != nil {
		return nil, fmt.Errorf("Could not convert to json: %v", err)
	}
	yamlbytes, err := yaml.JSONToYAML(jsonbytes)
	if err != nil {
		return nil, fmt.Errorf("Could not convert to yaml: %v", err)
	}
	err = os.WriteFile(outpath, yamlbytes, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("Could not write to %s: %v", outpath, err)
	}
	return topo, nil
}

func genTopo(ast *astopo.Topology) (*tpb.Topology, error) {
	ethCnt := make(map[int32]int32, 64)
	ethIPs := make(map[int32]map[int32]string, 64)
	ipCnt := make(map[int32]int32, 64)
	nodeMap := make(map[int32]*astopo.ASNode, 64)
	topo := &tpb.Topology{
		Name:  ast.Name,
		Nodes: make([]*tpb.Node, 0, 16),
		Links: make([]*tpb.Link, 0, 16),
	}

	for _, node := range ast.Nodes {
		ethCnt[node.Asn] = 0
		ethIPs[node.Asn] = make(map[int32]string, 4)
		nodeMap[node.Asn] = node
	}

	for _, link := range ast.Links {
		ethCnt[link.ANode]++
		ethCnt[link.ZNode]++
		ethCnt0 := ethCnt[link.ANode]
		ethCnt1 := ethCnt[link.ZNode]
		netSize0 := getSubnetSize(nodeMap[link.ANode].Net)
		netSize1 := getSubnetSize(nodeMap[link.ZNode].Net)
		tempLink := &tpb.Link{
			ANode: fmt.Sprintf("r%d", link.ANode),
			ZNode: fmt.Sprintf("r%d", link.ZNode),
			AInt:  fmt.Sprintf("eth%d", ethCnt0),
			ZInt:  fmt.Sprintf("eth%d", ethCnt1),
		}

		//Allocate IP to eth interfaces
		switch netSize0 >= netSize1 {
		case true:
			tempAInt, tempZInt := getAllocatedIP(nodeMap[link.ANode].Net, ipCnt[link.ANode])
			ethIPs[link.ANode][ethCnt0] = tempAInt
			ethIPs[link.ZNode][ethCnt1] = tempZInt
			ipCnt[link.ANode] += 4
		case false:
			tempAInt, tempZInt := getAllocatedIP(nodeMap[link.ZNode].Net, ipCnt[link.ZNode])
			ethIPs[link.ANode][ethCnt0] = tempAInt
			ethIPs[link.ZNode][ethCnt1] = tempZInt
			ipCnt[link.ZNode] += 4

		}
		topo.Links = append(topo.Links, tempLink)
	}
	for _, node := range ast.Nodes {
		tempnode := &tpb.Node{
			Name:   fmt.Sprintf("r%d", node.Asn),
			Type:   tpb.Type_BGPNODE,
			IpAddr: make(map[string]string, 4),
			Config: &tpb.Config{
				Tasks: []*tpb.Task{
					{Container: fmt.Sprintf("r%d", node.Asn), Cmds: []string{
						"/usr/local/bin/gobgpd -f /config/gobgp.toml > /dev/null 2> /dev/null &",
					}},
				},
				ExtraImages: map[string]string{
					fmt.Sprintf("r%d-frr", node.Asn): "frrouting/frr:v8.1.0",
				},
				ShareVolumes: []string{
					"zebra",
				},
				ContainerVolumes: map[string]*tpb.PublicVolumes{
					fmt.Sprintf("r%d-frr", node.Asn): {Volumes: []string{"zebra"}, Paths: []string{"/var/run/frr"}},
					fmt.Sprintf("r%d", node.Asn):     {Volumes: []string{"zebra"}, Paths: []string{"/var/run/frr"}},
				},
				// ConfigFile: fmt.Sprintf("r%d.toml", node.Asn),
			},
		}
		for eth, ip := range ethIPs[node.Asn] {
			tempnode.IpAddr[fmt.Sprintf("eth%d", eth)] = ip
		}
	}
	return topo, nil
}

func getSubnetSize(subnet string) int32 {
	strArr := strings.Split(subnet, "/")
	size, _ := strconv.Atoi(strArr[1])
	return int32(size)
}

func getAllocatedIP(subnet string, ipcnt int32) (string, string) {
	strArr := strings.Split(subnet, "/")
	ips := strings.Split(strArr[0], ".")
	ip3, _ := strconv.Atoi(ips[3])
	ips[3] = strconv.Itoa(ip3 + int(ipcnt) + 1)
	ipA := strings.Join(ips, ".")
	ips[3] = strconv.Itoa(ip3 + int(ipcnt) + 2)
	ipB := strings.Join(ips, ".")
	return ipA, ipB
}
