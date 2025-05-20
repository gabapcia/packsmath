package pack_test

import (
	"context"
	"errors"
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
		err := svc.RegisterPackSize(t.Context(), 1000)

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
		err := svc.RegisterPackSize(t.Context(), 500)

		require.NoError(t, err)
		assert.Equal(t, 1, len(mockStorage.RegisterPackSizeCalls()))
	})
}

func TestService_ListPackSizes(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		expected := []int{250, 500, 1000}

		mockStorage := &mock.PackStorageMock{
			ListPackSizesFunc: func(ctx context.Context) ([]int, error) {
				return expected, nil
			},
		}

		svc := pack.NewService(mockStorage)
		result, err := svc.ListPackSizes(context.Background())

		require.NoError(t, err)
		assert.Equal(t, expected, result)
		assert.Equal(t, 1, len(mockStorage.ListPackSizesCalls()))
	})

	t.Run("Returns Error From Storage", func(t *testing.T) {
		expectedErr := errors.New("storage failure")

		mockStorage := &mock.PackStorageMock{
			ListPackSizesFunc: func(ctx context.Context) ([]int, error) {
				return nil, expectedErr
			},
		}

		svc := pack.NewService(mockStorage)
		result, err := svc.ListPackSizes(context.Background())

		require.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedErr, err)
		assert.Equal(t, 1, len(mockStorage.ListPackSizesCalls()))
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
		err := svc.DeletePackSize(t.Context(), 2000)

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
		err := svc.DeletePackSize(t.Context(), 500)

		require.NoError(t, err)
		assert.Equal(t, 1, len(mockStorage.DeletePackSizeCalls()))
	})
}
