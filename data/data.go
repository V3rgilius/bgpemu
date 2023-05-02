package data

import (
	"context"
	"fmt"

	"github.com/p3rdy/bgpemu/lab"
	api "github.com/p3rdy/bgpemu/proto/gobgp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Start(topoName string) error {
	m, err := lab.New(topoName)
	if err != nil {
		return err
	}
	err = m.GetGrpcServersAll(topoName)
	if err != nil {
		return err
	}

	gServers := m.GetGServers()
	for router := range gServers {
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
		Filename:         fmt.Sprintf("%s-20060102.1504.dump", r),
		RotationInterval: 60,
	}
	resp, err := client.EnableMrt(context.Background(), req)
	if err != nil {
		return err
	}
	fmt.Printf("AddPath response: %s\n", resp.String())
	return nil
}
