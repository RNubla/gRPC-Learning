package main

import (
	"log"

	"github.com/RNubla/gRPC-Learning/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// addr := "localhost:9999"
	// conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer conn.Close()

	// // client := helloworldpb.NewHelloWorldClient(conn)
	// // req := helloworldpb.HelloWorldRequest{
	// // 	Request: "Hello World, This is gRPC working with go and python!",
	// // }
	// client := youtubeDL.NewGetVideoInfoServiceClient(conn)
	// req := youtubeDL.GetVideoInfoRequest{
	// 	VideoURL: "https://www.youtube.com/watch?v=8SnC354B5Vc",
	// }

	// resp, err := client.GetVideo(context.Background(), &req)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(toJSON(resp))

	app := fiber.New()
	router.SetupRoutes(app)
	log.Fatal(app.Listen("localhost:3000"))

}
