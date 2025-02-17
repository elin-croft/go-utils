package main

import (
	"flag"

	"github.com/elin-croft/go-utils/demo/api"
	"github.com/elin-croft/go-utils/framework/utils/startrek"
)

func main() {
	var rpcType string
	flag.StringVar(&rpcType, "type", "rpc", "rpc type")
	flag.Parse()

	var server startrek.Server
	if rpcType == "rpc" {
		server = api.NewDemoRpcServer()
	} else if rpcType == "grpc" {
		server = api.NewDemoGRPCServer()
	}
	server.Serve()
}
