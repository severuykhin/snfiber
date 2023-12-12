package snfiber

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func Send(c *fiber.Ctx, statusCode int, data interface{}) error {
	content, err := json.Marshal(data)

	if err != nil {
		return err
	}

	c.Set("Content-Type", "application/json")

	if err := c.SendStatus(statusCode); err != nil {
		return err
	}
	return c.Send(content)
}
