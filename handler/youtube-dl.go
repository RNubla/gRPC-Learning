package handler

import (
	"context"
	"log"

	"gRPC-Learning/gen/youtubeDL"

	"github.com/gofiber/fiber/v2"
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
func GetVideoInfo(c *fiber.Ctx) error {
	type vidUrl struct {
		VideoUrl string `json:videoUrl`
	}
	data := new(vidUrl)

	if err := c.BodyParser(data); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your indput", "data": err})
	}
	// fmt.Println(data)

	addr := "localhost:9999"
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := youtubeDL.NewGetVideoInfoServiceClient(conn)
	req := youtubeDL.GetVideoInfoRequest{
		VideoURL: data.VideoUrl,
	}

	resp, err := client.GetVideo(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}

	// if &resp.GetError().StatusCode == 404 {
	// 	return fiber.NewError(fiber.StatusNotFound, string("video not found"))
	// }

	return c.JSON(resp)

}
