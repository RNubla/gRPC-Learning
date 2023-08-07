package handler

import (
	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}

func TestPost(c *fiber.Ctx) error {
	type testPost struct {
		Message string `json:message`
	}

	data := new(testPost)

	if err := c.BodyParser(data); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your indput", "data": err})
	}

	return c.JSON(data)
}
