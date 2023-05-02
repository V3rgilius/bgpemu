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
	err := m.GetGrpcServers(pods, rd.TopoName)
	if err != nil {
		return err
	}
	for _, r := range rd.Routes {
		err := deployRoute(r)
		if err != nil {
			return err
		}
	}
	return nil
}

func deployRoute(r *rtpb.Route) error {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
func addPath(client api.GobgpApiClient, path *api.Path) error {
	// send the new route to the gobgp daemon
	req := &api.AddPathRequest{
		Path: path,
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
