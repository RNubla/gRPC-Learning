package main

import (
	"context"
	"log"

	"github.com/RNubla/gRPC-Learning/helloworldpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := "localhost:9999"
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := helloworldpb.NewHelloWorldClient(conn)
	req := helloworldpb.HelloWorldRequest{
		Request: "Hello World, This is gRPC working with go and python!",
	}

	resp, err := client.Print(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(resp.Response)
}
