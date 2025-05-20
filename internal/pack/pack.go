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
type PackStorage interface {
	// RegisterPackSize adds a new pack size
	RegisterPackSize(ctx context.Context, size int) error

	// DeletePackSize removes an existing pack size
	DeletePackSize(ctx context.Context, size int) error
}

// RegisterPackSize registers a new pack size via the underlying storage
func (s *service) RegisterPackSize(ctx context.Context, size int) error {
	return s.packStorage.RegisterPackSize(ctx, size)
}

// DeletePackSize deletes a pack size via the underlying storage
func (s *service) DeletePackSize(ctx context.Context, size int) error {
	return s.packStorage.DeletePackSize(ctx, size)
}
