package api

import (
	"fmt"

	"github.com/gabapcia/packsmath/internal/order"
	"github.com/gabapcia/packsmath/internal/pack"

	_ "github.com/gabapcia/packsmath/internal/handler/api/docs"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title PacksMath API
// @description An application that can calculate the number of packs we need to ship to the customer
func Run(port int, packService pack.Service, orderService order.Service) error {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ErrorHandler:          errorHandler,
	})

	app.Get("/docs/*", fiberSwagger.WrapHandler)

	app.Post("/packs", RegisterPackSizeHandler(packService))
	app.Get("/packs", ListPackSizesHandler(packService))
	app.Delete("/packs/:size", DeletePackSizeHandler(packService))

	app.Post("/orders", PackOrderHandler(orderService))

	return app.Listen(fmt.Sprintf(":%d", port))
}
