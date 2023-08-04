package main

import (
	"context"
	"fmt"
	"log"

	"github.com/RNubla/gRPC-Learning/gen/youtubeDL"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func toJSON(pb proto.Message) string {
	out, err := protojson.Marshal(pb)

	if err != nil {
		log.Fatalln("Cant convert to JSON", err)
	}

	return string(out)
}

func main() {
	addr := "localhost:9999"
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// client := helloworldpb.NewHelloWorldClient(conn)
	// req := helloworldpb.HelloWorldRequest{
	// 	Request: "Hello World, This is gRPC working with go and python!",
	// }
	client := youtubeDL.NewGetVideoInfoServiceClient(conn)
	req := youtubeDL.GetVideoInfoRequest{
		VideoURL: "https://www.youtube.com/watch?v=8SnC354B5Vc",
	}

	resp, err := client.GetVideo(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(toJSON(resp))

}
