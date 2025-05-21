package main

import (
	"github.com/gabapcia/packsmath/internal/handler/api"
	"github.com/gabapcia/packsmath/internal/infra/storage/memory"
	"github.com/gabapcia/packsmath/internal/order"
	"github.com/gabapcia/packsmath/internal/pack"
)

func main() {
	// Storages
	memoryStorage := memory.New()

	// Services
	packService := pack.NewService(memoryStorage)
	orderService := order.NewService(memoryStorage)

	// Handlers
	if err := api.Start(3000, packService, orderService); err != nil {
		panic(err)
	}
}
