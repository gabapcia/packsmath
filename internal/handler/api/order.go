package api

import (
	"github.com/gabapcia/packsmath/internal/order"

	"github.com/gofiber/fiber/v2"
)

type PackOrderRequest struct {
	Order uint64 `json:"order"`
}

type PackOrderResponse map[int]int

// PackOrder godoc
// @Summary      Pack An Order
// @Description  Packs an order
// @Accept       json
// @Param request body PackOrderRequest true "PackOrderRequest"
// @Success      200 {object} PackOrderResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /orders [post]
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
