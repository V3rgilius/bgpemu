package lab

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"

	api "github.com/p3rdy/bgpemu/proto/gobgp"
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
		fmt.Println(string(jsonBytes))
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
	err := m.GetGrpcServers(pods)
	if err != nil {
		return err
	}
	for _, pd := range pds.PolicyDeployments {
		err := deployPolicy(pd, m.GetGServers()[pd.RouterName])
		if err != nil {
			return err
		}
	}
	return nil
}

func deployPolicy(p *popb.PolicyDeployment, g string) error {
	conn, err := grpc.Dial(g, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	// create a client instance for the gRPC API
	client := api.NewGobgpApiClient(conn)
	defer conn.Close()
	err = addDefinedSets(client, p.DefinedSets)
	if err != nil {
		return err
	}
	log.Infof("DefinedSets added on: %s", p.RouterName)
	err = addStatements(client, p.Statements)
	if err != nil {
		return err
	}
	log.Infof("Statements added on: %s", p.RouterName)
	err = addPolicies(client, p.Policies)
	if err != nil {
		return err
	}
	log.Infof("Policies added on: %s", p.RouterName)
	// err = setPolicies(client, p)
	// if err != nil {
	// 	return err
	// }
	err = addPeerGroup(client, p.PeerGroups)
	if err != nil {
		return err
	}
	err = addPolicyAssignments(client, p.Assignments)
	if err != nil {
		return err
	}
	log.Infof("Assignments added on: %s", p.RouterName)
	return nil
}

func addStatements(client api.GobgpApiClient, ss []*api.Statement) error {
	for _, s := range ss {
		req := &api.AddStatementRequest{
			Statement: s,
		}
		_, err := client.AddStatement(context.Background(), req)
		if err != nil {
			return err
		}
	}
	return nil
}

func addDefinedSets(client api.GobgpApiClient, dss []*api.DefinedSet) error {
	for _, ds := range dss {
		req := &api.AddDefinedSetRequest{
			DefinedSet: ds,
		}
		_, err := client.AddDefinedSet(context.Background(), req)
		if err != nil {
			return err
		}
	}
	return nil
}

func addPolicies(client api.GobgpApiClient, ps []*api.Policy) error {
	for _, p := range ps {
		req := &api.AddPolicyRequest{
			Policy:                  p,
			ReferExistingStatements: true,
		}
		_, err := client.AddPolicy(context.Background(), req)
		if err != nil {
			return err
		}
	}
	return nil
}

func setPolicies(client api.GobgpApiClient, p *popb.PolicyDeployment) error {
	req := &api.SetPoliciesRequest{
		DefinedSets: p.DefinedSets,
		Policies:    p.Policies,
	}
	_, err := client.SetPolicies(context.Background(), req)
	if err != nil {
		return err
	}
	return nil
}

func addPolicyAssignments(client api.GobgpApiClient, pas []*api.PolicyAssignment) error {
	for _, pa := range pas {
		pa.Name = "global"
		req := &api.AddPolicyAssignmentRequest{
			Assignment: pa,
		}
		_, err := client.AddPolicyAssignment(context.Background(), req)
		if err != nil {
			return err
		}
	}
	return nil
}

func addPeerGroup(client api.GobgpApiClient, pgs []*api.PeerGroup) error {
	for _, pg := range pgs {
		req := &api.AddPeerGroupRequest{
			PeerGroup: pg,
		}
		_, err := client.AddPeerGroup(context.Background(), req)
		if err != nil {
			return err
		}
	}
	return nil
}
