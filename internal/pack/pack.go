package pack

import (
	"context"
	"errors"
)

var (
	// ErrPackSizeAlreadyRegistered indicates the pack size is already registered
	ErrPackSizeAlreadyRegistered = errors.New("pack size already registered")

	// ErrPackSizeNotFound indicates the pack size was not found
	ErrPackSizeNotFound = errors.New("pack size not found")
)

// PackStorage defines methods to register and delete pack sizes
//
//go:generate moq -pkg mock -out mock/pack_storage.go . PackStorage
type PackStorage interface {
	// RegisterPackSize adds a new pack size
	RegisterPackSize(ctx context.Context, size int) error

	// ListPackSizes returns the list of available pack sizes
	ListPackSizes(ctx context.Context) ([]int, error)

	// DeletePackSize removes an existing pack size
	DeletePackSize(ctx context.Context, size int) error
}

// RegisterPackSize registers a new pack size via the underlying storage
func (s *service) RegisterPackSize(ctx context.Context, size int) error {
	return s.packStorage.RegisterPackSize(ctx, size)
}

// ListPackSizes retrieves the list of available pack sizes from the underlying storage.
func (s *service) ListPackSizes(ctx context.Context) ([]int, error) {
	return s.packStorage.ListPackSizes(ctx)
}

// DeletePackSize deletes a pack size via the underlying storage
func (s *service) DeletePackSize(ctx context.Context, size int) error {
	return s.packStorage.DeletePackSize(ctx, size)
}
