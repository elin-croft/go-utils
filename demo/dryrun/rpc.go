package main

import (
	"context"
	"fmt"
	"log"
	"time"

	rpc_client "github.com/elin-croft/go-utils/framework/utils/startrek/thrift_rpc/client"
	"github.com/elin-croft/go-utils/gen-go/ew_go_thrift/ew_toolbox"
)

func main() {
	host := "localhost:9090"
	// serviceTransport := thrift.NewTSocketConf(host, &thrift.TConfiguration{
	// 	ConnectTimeout: 10 * time.Second,
	// 	SocketTimeout:  10 * time.Second,
	// })
	// if err := serviceTransport.Open(); err != nil {
	// 	log.Fatal("Error connecting to server:", err)
	// }
	// defer serviceTransport.Close()
	// protocol := thrift.NewTBinaryProtocolConf(serviceTransport, &thrift.TConfiguration{})
	// mq := thrift.NewTMultiplexedProtocol(protocol, "predict")
	// client := ew_toolbox.NewPredictClient(thrift.NewTStandardClient(mq, mq))
	t1 := time.Now().UnixMilli()
	client := rpc_client.NewClient(host, "predict", 10*time.Second)
	req := &ew_toolbox.Request{ID: 123}
	ctx := context.Background()
	conn := ew_toolbox.NewPredictClient(client)
	resp, err := conn.Get(ctx, req)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	fmt.Printf("res: %v\n", resp)
	fmt.Printf("takes %v millisecond\n", time.Now().UnixMilli()-t1)
}
