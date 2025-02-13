package server

import (
	"log"
	"time"

	"github.com/apache/thrift/lib/go/thrift"
)

func NewServer(serviceMap map[string]thrift.TProcessor, host string, timeout time.Duration) *thrift.TSimpleServer {
	processor := thrift.NewTMultiplexedProcessor()
	for name, service := range serviceMap {
		processor.RegisterProcessor(name, service)
	}
	serviceTransport, err := thrift.NewTServerSocketTimeout(host, timeout)
	if err != nil {
		log.Fatalf("Unable to create server socket: %v", err)
	}
	server := thrift.NewTSimpleServer2(processor, serviceTransport)
	return server
}
