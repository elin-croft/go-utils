package api

import (
	"fmt"
	"time"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/elin-croft/go-utils/demo/service"
	"github.com/elin-croft/go-utils/framework/utils/startrek"
	rpc_server "github.com/elin-croft/go-utils/framework/utils/startrek/thrift_rpc/server"
)

func NewDemoRpcServer() startrek.Server {
	serviceMap := map[string]thrift.TProcessor{
		"predict": service.NewProcessor(),
	}
	host := "localhost:9090"
	server := rpc_server.NewServer(serviceMap, host, 10*time.Second)
	fmt.Printf("server will start at: %s\n", host)
	return server
}
