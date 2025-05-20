package memory

import (
	"sync"

	"github.com/gabapcia/packsmath/internal/order"
	"github.com/gabapcia/packsmath/internal/pack"
)

// storage is a thread-safe, in-memory implementation of both order.PackStorage and pack.PackStorage. It stores pack sizes using a map for efficient lookup and modification
type storage struct {
	lock  sync.RWMutex
	packs map[int]struct{}
}

// New creates and returns a new in-memory pack storage instance
func New() *storage {
	return &storage{packs: make(map[int]struct{})}
}

// Ensure storage implements the pack.PackStorage interface
var _ pack.PackStorage = (*storage)(nil)

// Ensure storage implements the order.PackStorage interface
var _ order.PackStorage = (*storage)(nil)
