package client

import (
	"context"
	"log"
	"time"

	"github.com/apache/thrift/lib/go/thrift"
)

type Client struct {
	host           string
	serviceName    string
	TConfiguration *thrift.TConfiguration
}

func NewClient(host string, serviceName string, timeout time.Duration) *Client {
	opts := []Option{WithConnectionTimeout(timeout), WithSocketTimeout(timeout)}
	return NewClientWithOpt(host, serviceName, opts...)
}

func NewClientWithOpt(host string, serviceName string, opts ...Option) *Client {
	c := &Client{
		host:           host,
		serviceName:    serviceName,
		TConfiguration: &thrift.TConfiguration{},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func NewClientWithConfig(host string, serviceName string, config *thrift.TConfiguration) *Client {
	return &Client{
		host:           host,
		serviceName:    serviceName,
		TConfiguration: config,
	}
}

func (c *Client) Call(ctx context.Context, method string, args, result thrift.TStruct) (thrift.ResponseMeta, error) {
	serviceTransport := thrift.NewTSocketConf(c.host, c.TConfiguration)

	if err := serviceTransport.Open(); err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer serviceTransport.Close()

	protocol := thrift.NewTBinaryProtocolConf(serviceTransport, c.TConfiguration)
	mq := thrift.NewTMultiplexedProtocol(protocol, c.serviceName)

	conn := thrift.NewTStandardClient(mq, mq)
	resp, err := conn.Call(ctx, method, args, result)
	return resp, err
}
