package snfiber

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/severuykhin/goerrors"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	logger logger
}

type requestHandler func(req *Request) (interface{}, error)

func (h *handler) handle(handlerFunc requestHandler) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		req := Request{
			Ctx: c,
		}
		res, err := handlerFunc(&req)
		if err != nil {
			return h.handleError(c, err)
		}
		return Send(c, http.StatusOK, res)
	}
}

func (h *handler) handleError(c *fiber.Ctx, err error) error {

	appErr := goerrors.From(err)
	errKind := appErr.GetKind()
	errCtx := appErr.GetContext()
	errId := uuid.NewString()

	var statusCode int
	var message string

	switch {
	case errKind == goerrors.ErrBadRequest:
		statusCode = http.StatusBadRequest
		message = appErr.GetMessage()
	case errKind == goerrors.ErrNotFound:
		statusCode = http.StatusNotFound
		message = appErr.GetMessage()
	default:
		statusCode = http.StatusInternalServerError
		message = "internal server error"
	}

	logCtx := []interface{}{"id", errId, "kind", errKind, "statusCode", statusCode, "stacktrace", appErr.GetStackTrace(2)}
	logCtx = append(logCtx, errCtx.ToList()...)
	h.logger.Error(appErr.GetMessage(), logCtx...)

	responseData := map[string]interface{}{
		"error": map[string]interface{}{
			"message": message,
			"kind":    errKind,
			"id":      errId,
		},
	}

	return Send(c, statusCode, responseData)
}
