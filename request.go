package snfiber

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Request struct {
	*fiber.Ctx
}

func (c *Request) BodyParser(dest interface{}) error {
	body := c.Body()
	if len(body) == 0 {
		return fmt.Errorf("empty body")
	}

	err := json.Unmarshal(body, dest)
	if err != nil {
		switch err := err.(type) {
		case *json.UnmarshalTypeError:
			return fmt.Errorf("field:%s must be of type:%s instead has type:%s", err.Field, err.Type, err.Value)
		default:
			return err
		}
	}
	return nil
}
