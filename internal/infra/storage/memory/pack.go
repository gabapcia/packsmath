package memory

import (
	"context"

	"github.com/gabapcia/packsmath/internal/pack"
)

// RegisterPackSize adds a new pack size to the in-memory store
//
// Returns pack.ErrPackSizeAlreadyRegistered if the size is already present
func (s *storage) RegisterPackSize(ctx context.Context, size int) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, ok := s.packs[size]; ok {
		return pack.ErrPackSizeAlreadyRegistered
	}

	s.packs[size] = struct{}{}
	return nil
}

// ListPackSizes returns a slice of all registered pack sizes
func (s *storage) ListPackSizes(ctx context.Context) ([]int, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	packSizes := make([]int, 0, len(s.packs))
	for size := range s.packs {
		packSizes = append(packSizes, size)
	}

	return packSizes, nil
}

// DeletePackSize removes a pack size from the in-memory store
//
// Returns pack.ErrPackSizeNotFound if the size does not exist
func (s *storage) DeletePackSize(ctx context.Context, size int) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, ok := s.packs[size]; !ok {
		return pack.ErrPackSizeNotFound
	}

	delete(s.packs, size)
	return nil
}
