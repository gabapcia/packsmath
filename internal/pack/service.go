package pack

import "context"

// Service defines the operations available for managing pack sizes
//
//go:generate moq -pkg mock -out mock/service.go . Service
type Service interface {
	// RegisterPackSize registers a new pack size
	RegisterPackSize(ctx context.Context, size int) error

	// ListPackSizes returns the list of available pack sizes
	ListPackSizes(ctx context.Context) ([]int, error)

	// DeletePackSize deletes an existing pack size
	DeletePackSize(ctx context.Context, size int) error
}

// service implements the Service interface using a PackStorage backend
type service struct {
	packStorage PackStorage
}

// NewService creates a new Service using the given PackStorage
func NewService(packStorage PackStorage) Service {
	return &service{
		packStorage: packStorage,
	}
}
