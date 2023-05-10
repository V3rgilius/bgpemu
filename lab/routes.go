package lab

import (
	"context"
	"fmt"
	"os"
	"strings"

	"google.golang.org/grpc/credentials/insecure"

	api "github.com/p3rdy/bgpemu/proto/gobgp"
	rtpb "github.com/p3rdy/bgpemu/proto/routes"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"sigs.k8s.io/yaml"
)

// LoadRoutes loads a Topology from path and parse all subtopos.
func LoadRoutes(path string) (*rtpb.RouteDeployment, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	rd := &rtpb.RouteDeployment{}
	switch {
	case strings.HasSuffix(path, ".yaml"), strings.HasSuffix(path, ".yml"):
		jsonBytes, err := yaml.YAMLToJSON(b)
		if err != nil {
			return nil, fmt.Errorf("could not parse yaml: %v", err)
		}
		if err := protojsonUnmarshaller.Unmarshal(jsonBytes, rd); err != nil {
			return nil, fmt.Errorf("could not parse json: %v", err)
		}
	default:
		if err := prototext.Unmarshal(b, rd); err != nil {
			return nil, err
		}
	}
	return rd, nil
}

func (m *Manager) DeployRoutes(rd *rtpb.RouteDeployment) error {
	// 获取路由文件中设备对应的Pod的gRPC接口
	// 创建gRPC连接
	// 构造，调用
	pods := make([]string, 0, 16)
	for _, r := range rd.Routes {
		pods = append(pods, r.Name)
	}
	err := m.GetGrpcServers(pods)
	if err != nil {
		return err
	}
	for _, r := range rd.Routes {
		err := deployRoute(r, m.GetGServers()[r.Name])
		if err != nil {
			return err
		}
	}
	return nil
}

func deployRoute(r *rtpb.Route, g string) error {
	conn, err := grpc.Dial(g, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	// create a client instance for the gRPC API
	client := api.NewGobgpApiClient(conn)
	defer conn.Close()

	for _, ri := range r.Paths {
		err := addPath(client, ri)
		if err != nil {
			return err
		}
	}
	// print the response from the gobgp daemon
	return nil
}
func addPath(client api.GobgpApiClient, path *rtpb.BgpPath) error {
	// send the new route to the gobgp daemon
	nlriAny := &anypb.Any{}
	err := anypb.MarshalFrom(nlriAny, path.Nlri, proto.MarshalOptions{})
	if err != nil {
		return err
	}
	nextHop := &anypb.Any{}
	err = anypb.MarshalFrom(nextHop, &api.NextHopAttribute{
		NextHop: "0.0.0.0",
	}, proto.MarshalOptions{})
	if err != nil {
		return err
	}
	origin := &anypb.Any{}
	err = anypb.MarshalFrom(origin, &api.OriginAttribute{
		Origin: 0,
	}, proto.MarshalOptions{})
	if err != nil {
		return err
	}
	req := &api.AddPathRequest{
		Path: &api.Path{
			Nlri: nlriAny,
			Family: &api.Family{
				Afi:  api.Family_AFI_IP,
				Safi: api.Family_SAFI_UNICAST,
			},
			Pattrs: []*anypb.Any{nextHop, origin},
		},
	}
	resp, err := client.AddPath(context.Background(), req)
	if err != nil {
		return err
	}
	fmt.Printf("AddPath response: %s\n", resp.String())
	return nil
}

// func parsePrefix(addr string) *api.IPAddressPrefix {
// 	splitArr := strings.Split(addr, "/")
// 	len, _ := strconv.Atoi(splitArr[1])
// 	return &api.IPAddressPrefix{
// 		Prefix:    splitArr[0],
// 		PrefixLen: uint32(len),
// 	}
// }
