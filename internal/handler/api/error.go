package api

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/gabapcia/packsmath/internal/pack"

	"github.com/gofiber/fiber/v2"
)

// ErrorResponse represents the structure of an error returned to the client.
type ErrorResponse struct {
	Message string `json:"message"`
}

var (
	ErrorResponseInvalidRequestBody = ErrorResponse{Message: "invalid request body"}
	ErrorResponseUnknownError       = ErrorResponse{Message: "unknown error"}
)

// ErrorHandler maps known application errors to appropriate HTTP status codes and messages.
// If the error is not recognized, it logs the error and returns a generic 500 response
//
// This function is designed to be used as Fiber's centralized error handler
func ErrorHandler(c *fiber.Ctx, err error) error {
	var syntaxErr *json.SyntaxError

	switch {
	case errors.Is(err, pack.ErrPackSizeAlreadyRegistered):
		return c.Status(http.StatusConflict).JSON(ErrorResponse{Message: err.Error()})
	case errors.Is(err, pack.ErrPackSizeNotFound):
		return c.Status(http.StatusNotFound).JSON(ErrorResponse{Message: err.Error()})
	case errors.As(err, &syntaxErr):
		return c.Status(http.StatusBadRequest).JSON(ErrorResponseInvalidRequestBody)
	default:
		slog.Error("unknown error", "error", err)
		return c.Status(http.StatusInternalServerError).JSON(ErrorResponseUnknownError)
	}
}
