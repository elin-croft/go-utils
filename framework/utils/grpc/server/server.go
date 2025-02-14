package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	host     string
	server   *grpc.Server
	listener net.Listener
}

func NewGRPCServer(host string, server *grpc.Server) *GRPCServer {
	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
		return nil
	}
	return &GRPCServer{
		host:     host,
		server:   server,
		listener: listener,
	}
}

func (s *GRPCServer) Serve() error {
	fmt.Printf("gRPC server is running on %v\n", s.host)
	if err := s.server.Serve(s.listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
		return err
	}
	return nil
}
