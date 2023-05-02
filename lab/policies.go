package lab

import (
	"fmt"
	"os"
	"strings"

	api "github.com/p3rdy/bgpemu/proto/gobgp"
	apipb "github.com/p3rdy/bgpemu/proto/gobgp"
	popb "github.com/p3rdy/bgpemu/proto/policies"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/prototext"
	"sigs.k8s.io/yaml"
)

// LoadRoutes loads a Topology from path and parse all subtopos.
func LoadPolicies(path string) (*popb.PolicyDeployments, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	pds := &popb.PolicyDeployments{}
	switch {
	case strings.HasSuffix(path, ".yaml"), strings.HasSuffix(path, ".yml"):
		jsonBytes, err := yaml.YAMLToJSON(b)
		if err != nil {
			return nil, fmt.Errorf("could not parse yaml: %v", err)
		}
		if err := protojsonUnmarshaller.Unmarshal(jsonBytes, pds); err != nil {
			return nil, fmt.Errorf("could not parse json: %v", err)
		}
	default:
		if err := prototext.Unmarshal(b, pds); err != nil {
			return nil, err
		}
	}
	return pds, nil
}

func (m *Manager) DeployPolicies(pds *popb.PolicyDeployments) error {
	// 获取路由文件中设备对应的Pod的gRPC接口
	// 创建gRPC连接
	// 构造，调用
	pods := make([]string, 0, 16)
	for _, pd := range pds.PolicyDeployments {
		pods = append(pods, pd.RouterName)
	}
	err := m.GetGrpcServers(pods, pds.TopoName)
	if err != nil {
		return err
	}
	for _, pd := range pds.PolicyDeployments {
		err := deployPolicy(pd)
		if err != nil {
			return err
		}
	}
	return nil
}

func deployPolicy(p *popb.PolicyDeployment) error {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	// create a client instance for the gRPC API
	client := api.NewGobgpApiClient(conn)
	defer conn.Close()

	err = setPolicies(client, p)
	if err != nil {
		return err
	}
	err = addPeerGroup(client, p.PeerGroups)
	if err != nil {
		return err
	}
	err = addPolicyAssignment(client, p.Assignments)
	if err != nil {
		return err
	}
	return nil
}

func setPolicies(client apipb.GobgpApiClient, p *popb.PolicyDeployment) error {

	return nil
}

func addPolicyAssignment(client apipb.GobgpApiClient, pa []*apipb.PolicyAssignment) error {
	return nil
}

func addPeerGroup(client apipb.GobgpApiClient, pg []*apipb.PeerGroup) error {
	return nil
}
