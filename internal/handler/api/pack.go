package api

import (
	"net/http"

	"github.com/gabapcia/packsmath/internal/pack"

	"github.com/gofiber/fiber/v2"
)

type RegisterPackRequest struct {
	Size uint64 `json:"size"`
}

// RegisterPackSizeHandler godoc
// @Summary      Register Pack Size
// @Description  Register a new pack size
// @Accept       json
// @Param request body RegisterPackRequest true "RegisterPackRequest"
// @Success      204
// @Failure      409  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /packs [post]
func RegisterPackSizeHandler(packService pack.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req RegisterPackRequest
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		if err := packService.RegisterPackSize(c.Context(), int(req.Size)); err != nil {
			return err
		}

		return c.SendStatus(http.StatusNoContent)
	}
}

// ListPackSizesHandler godoc
// @Summary      List Pack Sizes
// @Description  List all pack sizes
// @Accept       json
// @Produce      json
// @Success      200 {array} int
// @Failure      500  {object}  ErrorResponse
// @Router       /packs [get]
func ListPackSizesHandler(packService pack.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		packSizes, err := packService.ListPackSizes(c.Context())
		if err != nil {
			return err
		}

		return c.JSON(packSizes)
	}
}

type DeletePackSizeRequest struct {
	Size uint64 `params:"size"`
}

// DeletePackSizeHandler godoc
// @Summary      Delete Pack Size
// @Description  Deletes a pack size
// @Accept       json
// @Param        size   path      int  true  "pack size"
// @Success      204
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /packs/{size} [delete]
func DeletePackSizeHandler(packService pack.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req DeletePackSizeRequest
		if err := c.ParamsParser(&req); err != nil {
			return err
		}

		if err := packService.DeletePackSize(c.Context(), int(req.Size)); err != nil {
			return err
		}

		return c.SendStatus(http.StatusNoContent)
	}
}
