package client

import (
	"os"

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

type svrClient struct {
	c           client.Client
	serviceName string
}

func init() {
	// setup broker/client/server
	broker.DefaultBroker = bkr.NewBroker()
	client.DefaultClient = cli.NewClient()
	server.DefaultServer = srv.NewServer()
	cmd.Init()
}
// NewClient 新建客户连接
func NewClient(serviceName string) OrderPayClient {
	return &svrClient{
		c:           client.DefaultClient(),
		serviceName: serviceName,
	}
}