package router

import (
	"github.com/RNubla/gRPC-Learning/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	// YTDL GetInfo
	ytdl := api.Group("/ytdl")
	ytdl.Post("/", handler.GetVideoInfo)
}
