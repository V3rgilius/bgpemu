package data

import (
	"context"
	"fmt"
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/v3rgilius/bgpemu/lab"
	api "github.com/v3rgilius/bgpemu/proto/gobgp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Start(topoName string) error {
	m, err := lab.New(topoName)
	if err != nil {
		return err
	}
	err = m.GetGrpcServersAll()
	if err != nil {
		return err
	}

	gServers := m.GetGServers()
	for router := range gServers {
		log.Infof("Start on %s at %s", router, gServers[router])
		err = startMRT(router, gServers[router])
		if err != nil {
			return err
		}
	}
	return nil
}

func startMRT(r string, g string) error {
	conn, err := grpc.Dial(g, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	// create a client instance for the gRPC API
	client := api.NewGobgpApiClient(conn)
	defer conn.Close()
	req := &api.EnableMrtRequest{
		Type:             api.EnableMrtRequest_UPDATES,
		Filename:         "/tmp/log/20060102.1504.updates.dump",
		RotationInterval: 240,
	}
	_, err = client.EnableMrt(context.Background(), req)
	if err != nil {
		return err
	}
	// req = &api.EnableMrtRequest{
	// 	Type:             api.EnableMrtRequest_TABLE,
	// 	Filename:         "/tmp/log/20060102.1504.table.dump",
	// 	RotationInterval: 60,
	// }
	// _, err = client.EnableMrt(context.Background(), req)
	// if err != nil {
	// 	return err
	// }
	return nil
}

func Stop(topoName string) error {
	m, err := lab.New(topoName)
	if err != nil {
		return err
	}
	err = m.GetGrpcServersAll()
	if err != nil {
		return err
	}

	gServers := m.GetGServers()
	for router := range gServers {
		log.Infof("Stop on %s at %s", router, gServers[router])
		err = stopMRT(router, gServers[router])
		if err != nil {
			return err
		}
	}
	return nil
}

func stopMRT(r string, g string) error {
	conn, err := grpc.Dial(g, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	// create a client instance for the gRPC API
	client := api.NewGobgpApiClient(conn)
	defer conn.Close()
	log.Infof("disabling1")
	req := &api.DisableMrtRequest{
		Filename: "/tmp/log/20060102.1504.updates.dump",
	}
	_, err = client.DisableMrt(context.Background(), req)
	if err != nil {
		return err
	}
	log.Infof("disabling2")
	req = &api.DisableMrtRequest{
		Filename: "/tmp/log/20060102.1504.table.dump",
	}
	_, err = client.DisableMrt(context.Background(), req)
	if err != nil {
		return err
	}
	return nil
}

func Dump(topoName string) error {
	m, err := lab.New(topoName)
	if err != nil {
		return err
	}
	err = m.GetGrpcServersAll()
	if err != nil {
		return err
	}

	gServers := m.GetGServers()
	for router := range gServers {
		args := []string{
			"cp",
			fmt.Sprintf("%s/%s:tmp/log", topoName, router),
			fmt.Sprintf("./mrts/%s/", router),
		}
		pycmd := exec.Command("kubectl", args...)
		out, err := pycmd.Output() // 执行命令，并获取输出和错误信息
		if err != nil {
			return err
		}
		fmt.Println(string(out))
	}
	return nil
}
