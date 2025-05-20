package pack_test

import (
	"context"
	"testing"

	"github.com/gabapcia/packsmath/internal/pack"
	"github.com/gabapcia/packsmath/internal/pack/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_RegisterPackSize(t *testing.T) {
	t.Run("Duplicate Pack Size Returns Error", func(t *testing.T) {
		mockStorage := &mock.PackStorageMock{
			RegisterPackSizeFunc: func(ctx context.Context, size int) error {
				assert.Equal(t, 1000, size)
				return pack.ErrPackSizeAlreadyRegistered
			},
		}

		svc := pack.NewService(mockStorage)
		err := svc.RegisterPackSize(context.Background(), 1000)

		require.Error(t, err)
		assert.Equal(t, pack.ErrPackSizeAlreadyRegistered, err)
		assert.Equal(t, 1, len(mockStorage.RegisterPackSizeCalls()))
	})

	t.Run("Success", func(t *testing.T) {
		mockStorage := &mock.PackStorageMock{
			RegisterPackSizeFunc: func(ctx context.Context, size int) error {
				assert.Equal(t, 500, size)
				return nil
			},
		}

		svc := pack.NewService(mockStorage)
		err := svc.RegisterPackSize(context.Background(), 500)

		require.NoError(t, err)
		assert.Equal(t, 1, len(mockStorage.RegisterPackSizeCalls()))
	})
}

func TestService_DeletePackSize(t *testing.T) {
	t.Run("Delete Non Existent Pack Size Returns Error", func(t *testing.T) {
		t.Parallel()

		mockStorage := &mock.PackStorageMock{
			DeletePackSizeFunc: func(ctx context.Context, size int) error {
				assert.Equal(t, 2000, size)
				return pack.ErrPackSizeNotFound
			},
		}

		svc := pack.NewService(mockStorage)
		err := svc.DeletePackSize(context.Background(), 2000)

		require.Error(t, err)
		assert.Equal(t, pack.ErrPackSizeNotFound, err)
		assert.Equal(t, 1, len(mockStorage.DeletePackSizeCalls()))
	})

	t.Run("Success", func(t *testing.T) {
		mockStorage := &mock.PackStorageMock{
			DeletePackSizeFunc: func(ctx context.Context, size int) error {
				assert.Equal(t, 500, size)
				return nil
			},
		}

		svc := pack.NewService(mockStorage)
		err := svc.DeletePackSize(context.Background(), 500)

		require.NoError(t, err)
		assert.Equal(t, 1, len(mockStorage.DeletePackSizeCalls()))
	})
}
