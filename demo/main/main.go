package main

import (
	"fmt"
	"time"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/elin-croft/go-utils/demo/service"
	"github.com/elin-croft/go-utils/framework/utils/thrift_rpc/server"
)

func main() {
	serviceMap := map[string]thrift.TProcessor{
		"predict": service.NewProcessor(),
	}
	host := "localhost:9090"
	server := server.NewServer(serviceMap, host, 10*time.Second)
	fmt.Printf("server started at: %s\n", host)
	server.Serve()
}
