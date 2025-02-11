package client

import (
	"context"
	"log"

	"github.com/apache/thrift/lib/go/thrift"
)

type Client struct {
	host           string
	TConfiguration *thrift.TConfiguration
}

func NewClient(host string) *Client {
	return &Client{
		host:           host,
		TConfiguration: &thrift.TConfiguration{},
	}
}

func NewClientWithOpt(host string, opts ...Option) *Client {
	c := &Client{
		host:           host,
		TConfiguration: &thrift.TConfiguration{},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func NewClientWithConfig(host string, config *thrift.TConfiguration) *Client {
	return &Client{
		host:           host,
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
	mq := thrift.NewTMultiplexedProtocol(protocol, "predict")

	conn := thrift.NewTStandardClient(mq, mq)
	resp, err := conn.Call(ctx, method, args, result)
	return resp, err
}
