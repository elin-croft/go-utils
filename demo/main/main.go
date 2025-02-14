package main

import (
	"flag"

	"github.com/elin-croft/go-utils/demo/api"
)

type RpcServer interface {
	Serve() error
}

func main() {
	var rpcType string
	flag.StringVar(&rpcType, "type", "rpc", "rpc type")
	flag.Parse()

	var server RpcServer
	if rpcType == "rpc" {
		server = api.NewDemoRpcServer()
	} else if rpcType == "grpc" {
		server = api.NewDemoGRPCServer()
	}
	server.Serve()
}
