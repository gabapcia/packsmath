package api

import (
	"fmt"

	"github.com/gabapcia/packsmath/internal/order"
	"github.com/gabapcia/packsmath/internal/pack"

	"github.com/gofiber/fiber/v2"
)

// @title PacksMath API
// @description An application that can calculate the number of packs we need to ship to the customer
func Start(port int, packService pack.Service, orderService order.Service) error {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ErrorHandler:          errorHandler,
	})

	app.Post("/packs", RegisterPackSizeHandler(packService))
	app.Get("/packs", ListPackSizesHandler(packService))
	app.Delete("/packs", DeletePackSizeHandler(packService))

	return app.Listen(fmt.Sprintf(":%d", port))
}
