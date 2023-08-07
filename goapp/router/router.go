package router

import (
	"gapp/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	hello := app.Group("/hello", logger.New())
	hello.Get("/", handler.Hello)

	testPos := app.Group("/test-post", logger.New())
	testPos.Post("/", handler.TestPost)
}
