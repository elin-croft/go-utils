package main

import (
	"context"
	"fmt"
	"log"
	"time"

	grpc_client "github.com/elin-croft/go-utils/framework/utils/startrek/grpc/client"
	"github.com/elin-croft/go-utils/gen-go/protos/servers"
)

func main() {
	host := "localhost:9999"
	conn, err := grpc_client.NewClient(host)
	if err != nil {
		log.Fatalf("conn err: %v\n", conn)
		return
	}

	t1 := time.Now().UnixMilli()
	cli := servers.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &servers.HelloRequest{
		Name: "test",
	}
	resp, err := cli.SayHello(ctx, req)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	fmt.Printf("response: %v\n", resp)
	fmt.Printf("takes %v millisecond\n", time.Now().UnixMilli()-t1)
}
