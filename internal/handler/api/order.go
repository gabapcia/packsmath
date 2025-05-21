package api

import (
	"github.com/gabapcia/packsmath/internal/order"

	"github.com/gofiber/fiber/v2"
)

// PackOrderRequest represents the JSON body sent to the /orders endpoint
type PackOrderRequest struct {
	Order uint64 `json:"order"` // Quantity of items requested in the order
}

// PackOrderResponse represents the response map of pack sizes to quantities returned by the service.
type PackOrderResponse map[int]int

// PackOrderHandler returns a Fiber handler that processes POST requests to pack an order.
//
// It expects a JSON body with an `order` field and returns a map of pack sizes to quantities.
// In case of parsing or service errors, the handler returns the appropriate HTTP status codes.
//
// Swagger annotations:
//
//	@Summary      Pack An Order
//	@Description  Packs an order
//	@Accept       json
//	@Param        request body PackOrderRequest true "PackOrderRequest"
//	@Success      200 {object} PackOrderResponse
//	@Failure      500 {object} ErrorResponse
//	@Router       /orders [post]
func PackOrderHandler(orderService order.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req PackOrderRequest
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		packs, err := orderService.PackOrder(c.Context(), int(req.Order))
		if err != nil {
			return err
		}

		return c.JSON(packs)
	}
}
