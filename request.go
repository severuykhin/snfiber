package snfiber

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/severuykhin/goerrors"
)

type Request struct {
	*fiber.Ctx
}

func (c *Request) BodyParser(dest interface{}) error {
	body := c.Body()
	if len(body) == 0 {
		return goerrors.NewBadRequestErr().WithMessage("empty body")
	}

	err := json.Unmarshal(body, dest)
	if err != nil {
		switch err := err.(type) {
		case *json.UnmarshalTypeError:
			return goerrors.NewBadRequestErr().
				WithMessage(fmt.Sprintf("field:%s must be of type:%s instead has type:%s", err.Field, err.Type, err.Value))
		default:
			return goerrors.NewBadRequestErr().WithMessage(err.Error())
		}
	}
	return nil
}
