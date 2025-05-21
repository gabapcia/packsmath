package api

import (
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

// errorHandler maps known application errors to appropriate HTTP status codes and messages.
// If the error is not recognized, it logs the error and returns a generic 500 response
//
// This function is designed to be used as Fiber's centralized error handler
func errorHandler(c *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, pack.ErrPackSizeAlreadyRegistered):
		return c.Status(http.StatusConflict).JSON(ErrorResponse{Message: err.Error()})
	case errors.Is(err, pack.ErrPackSizeNotFound):
		return c.Status(http.StatusNotFound).JSON(ErrorResponse{Message: err.Error()})
	default:
		slog.Error("unknown error", "error", err)
		return c.Status(http.StatusInternalServerError).JSON(ErrorResponse{Message: "unknown error"})
	}
}
