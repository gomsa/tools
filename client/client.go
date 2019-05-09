package client

import (
	context "golang.org/x/net/context"

	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
	bkr "github.com/micro/go-plugins/broker/grpc"
	cli "github.com/micro/go-plugins/client/grpc"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	srv "github.com/micro/go-plugins/server/grpc"

	// static selector offloads load balancing to k8s services
	// enable with MICRO_SELECTOR=static or --selector=static
	// requires user to create k8s services
	_ "github.com/micro/go-plugins/selector/static"

)

func init() {
	// setup broker/client/server
	broker.DefaultBroker = bkr.NewBroker()
	client.DefaultClient = cli.NewClient()
	server.DefaultServer = srv.NewServer()
	cmd.Init()
}

// Request 请求
func Request(ctx context.Context, serviceName string, method string, in *interface{}, out *interface{}, opts ...client.CallOption) error {
	c := client.DefaultClient
	req := c.NewRequest(serviceName, method, in)
	err := c.Call(ctx, req, out, opts...)
	if err != nil {
		return  err
	}
	return  nil
}