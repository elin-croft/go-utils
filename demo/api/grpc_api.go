package api

import (
	"context"
	"fmt"

	grpc_server "github.com/elin-croft/go-utils/framework/utils/startrek/grpc/server"
	"github.com/elin-croft/go-utils/gen-go/protos/servers"
	"google.golang.org/grpc"
)

type Greeter struct {
	servers.UnimplementedGreeterServer
}

var _ servers.GreeterServer = (*Greeter)(nil)

func (s *Greeter) SayHello(ctx context.Context, req *servers.HelloRequest) (*servers.HelloResponse, error) {
	fmt.Printf("req: %+v\n", req)
	resp := &servers.HelloResponse{
		Message: "hello world",
	}
	return resp, nil
}

func NewDemoGRPCServer() *grpc_server.GRPCServer {
	host := "localhost:9999"
	gServer := grpc.NewServer()
	servers.RegisterGreeterServer(gServer, &Greeter{})
	return grpc_server.NewGRPCServer(host, gServer)
}
