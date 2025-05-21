package api

import (
	"fmt"

	"github.com/gabapcia/packsmath/internal/order"
	"github.com/gabapcia/packsmath/internal/pack"

	_ "github.com/gabapcia/packsmath/internal/handler/api/docs" // Swagger docs
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title           PacksMath API
// @description     An application that can calculate the number of packs we need to ship to the customer

// Run initializes the HTTP server and registers all API routes
//
// It accepts service implementations for pack and order operations, configures error handling,
// mounts Swagger documentation at `/docs/*`, and starts the server on the specified port
func Run(port int, packService pack.Service, orderService order.Service) error {
	app := fiber.New(fiber.Config{
		ErrorHandler: ErrorHandler,
	})

	// Swagger docs endpoint
	app.Get("/docs/*", fiberSwagger.WrapHandler)

	// Pack management endpoints
	app.Post("/packs", RegisterPackSizeHandler(packService))
	app.Get("/packs", ListPackSizesHandler(packService))
	app.Delete("/packs/:size", DeletePackSizeHandler(packService))

	// Order packing endpoint
	app.Post("/orders", PackOrderHandler(orderService))

	// Start HTTP server
	return app.Listen(fmt.Sprintf(":%d", port))
}
