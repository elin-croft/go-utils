package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func getDefaultOpt() []grpc.DialOption {
	withInsecure := grpc.WithTransportCredentials(insecure.NewCredentials())
	return []grpc.DialOption{withInsecure}
}
func NewClient(host string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, getDefaultOpt()...)
	conn, err := grpc.NewClient(host, opts...)
	return conn, err
}
