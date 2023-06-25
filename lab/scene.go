package lab

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	api "github.com/p3rdy/bgpemu/proto/gobgp"
	log "github.com/sirupsen/logrus"

	spb "github.com/p3rdy/bgpemu/proto/scene"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/prototext"
	"sigs.k8s.io/yaml"
)

func LoadScene(path string) (*spb.Scene, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	pds := &spb.Scene{}
	switch {
	case strings.HasSuffix(path, ".yaml"), strings.HasSuffix(path, ".yml"):
		jsonBytes, err := yaml.YAMLToJSON(b)
		if err != nil {
			return nil, fmt.Errorf("could not parse yaml: %v", err)
		}
		// fmt.Println(string(jsonBytes))
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
func DeployScene(scene *spb.Scene) error {
	m, err := New(scene.TopoName)
	if err != nil {
		return err
	}
	err = m.GetGrpcServersAll()
	if err != nil {
		return err
	}
	for _, b := range scene.Inits {
		err := behave(b, m)
		if err != nil {
			return err
		}
	}
	if scene.PoliciesPath != "" {
		pds, err := LoadPolicies(scene.PoliciesPath)
		if err != nil {
			return err
		}
		err = m.DeployPolicies(pds)
		if err != nil {
			return err
		}
	}
	if scene.RoutesPath != "" {
		rts, err := LoadRoutes(scene.RoutesPath)
		if err != nil {
			return err
		}
		err = m.DeployRoutes(rts)
		if err != nil {
			return err
		}
	}
	for _, b := range scene.Behaviors {
		err := behave(b, m)
		if err != nil {
			return err
		}
	}
	return nil
}
func behave(b *spb.Behavior, m *Manager) error {
	log.Infof("Processing behavior: %s", b.Name)
	for _, step := range b.Steps {
		err := execStep(step, b.DeviceName, m)
		if err != nil {
			return err
		}
	}
	return nil
}

func execStep(step *spb.Step, dn string, m *Manager) error {
	log.Infof("Executing step: %s", step.Name)
	switch body := step.Body.(type) {
	case *spb.Step_Cmds:
		cmds := body.Cmds
		err := execCmds(cmds, dn, m)
		if err != nil {
			return err
		}
	case *spb.Step_Aps:
		ap := body.Aps
		err := addPeers(ap.Peers, m.GetGServers()[dn])
		if err != nil {
			return err
		}
	case *spb.Step_Sbs:
		sb := body.Sbs
		err := startBgp(sb, m.GetGServers()[dn])
		if err != nil {
			return err
		}
	default:
		return nil
	}
	return nil
}

func execCmds(cmds *spb.Commands, deviceName string, m *Manager) error {
	err := m.Exec(context.Background(), cmds.Cmds, deviceName, cmds.Container, nil, os.Stdout, os.Stderr)
	return err
}

func transportFile(file *spb.FileTrans) error {
	return nil
}

func wait(t *spb.Wait) {
	// timenow := time.Now().Unix()
	if t.Timestamp != 0 {
		time.Until(time.Unix(int64(t.Timestamp), 0))
	}
}

func addPeers(peers []*api.Peer, g string) error {

	conn, err := grpc.Dial(g, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	// create a client instance for the gRPC API
	client := api.NewGobgpApiClient(conn)
	defer conn.Close()
	for _, p := range peers {
		req := &api.AddPeerRequest{
			Peer: p,
		}
		_, err = client.AddPeer(context.Background(), req)
		if err != nil {
			return err
		}
	}
	return nil
}

func startBgp(sbs *spb.StartBgpStep, g string) error {
	conn, err := grpc.Dial(g, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	client := api.NewGobgpApiClient(conn)
	defer conn.Close()
	startReq := &api.StartBgpRequest{
		Global: sbs.Global,
	}
	// create a client instance for the gRPC API
	_, err = client.StartBgp(context.Background(), startReq)
	if err != nil {
		return err
	}
	zebraReq := &api.EnableZebraRequest{
		Url:        "unix:/var/run/frr/zserv.api",
		RouteTypes: []string{},
		Version:    6,
	}
	// create a client instance for the gRPC API
	_, err = client.EnableZebra(context.Background(), zebraReq)
	if err != nil {
		return err
	}

	if sbs.GetRpki() != nil {
		_, err = client.AddRpki(context.Background(), sbs.GetRpki())
		if err != nil {
			return err
		}
	}
	return nil
}
