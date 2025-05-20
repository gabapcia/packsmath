package order

import "context"

// Service defines the operations for resolving how to pack customer orders
//
//go:generate moq -pkg mock -out mock/service.go . Service
type Service interface {
	// PackOrder calculates the pack distribution needed to fulfill the given order quantity.
	// It returns a map of pack size to quantity and any error encountered
	PackOrder(ctx context.Context, order int) (map[int]int, error)
}

// service is the concrete implementation of the Service interface
type service struct {
	packStorage PackStorage
}

// NewService creates a new instance of the order service using the provided PackStorage
func NewService(packStorage PackStorage) *service {
	return &service{
		packStorage: packStorage,
	}
}
