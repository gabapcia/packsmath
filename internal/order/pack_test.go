package order_test

import (
	"context"
	"errors"
	"testing"

	"github.com/gabapcia/packsmath/internal/order"
	"github.com/gabapcia/packsmath/internal/order/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_PackOrder(t *testing.T) {
	t.Run("Storage Returns Rrror", func(t *testing.T) {
		expectedErr := errors.New("storage failure")
		mockStorage := &mock.PackStorageMock{
			ListPackSizesFunc: func(ctx context.Context) ([]int, error) {
				return nil, expectedErr
			},
		}

		svc := order.NewService(mockStorage)
		result, err := svc.PackOrder(t.Context(), 500)

		require.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedErr, err)
		assert.Equal(t, 1, len(mockStorage.ListPackSizesCalls()))
	})

	t.Run("Fallback Case From Resolver Is Used", func(t *testing.T) {
		mockStorage := &mock.PackStorageMock{
			ListPackSizesFunc: func(ctx context.Context) ([]int, error) {
				return []int{100, 200, 500}, nil
			},
		}

		svc := order.NewService(mockStorage)
		result, err := svc.PackOrder(t.Context(), 499)

		require.NoError(t, err)
		assert.Equal(t, map[int]int{500: 1}, result)
	})

	t.Run("Success", func(t *testing.T) {
		mockStorage := &mock.PackStorageMock{
			ListPackSizesFunc: func(ctx context.Context) ([]int, error) {
				return []int{250, 500, 1000}, nil
			},
		}

		svc := order.NewService(mockStorage)
		result, err := svc.PackOrder(t.Context(), 1250)

		require.NoError(t, err)
		assert.Equal(t, map[int]int{1000: 1, 250: 1}, result)
		assert.Equal(t, 1, len(mockStorage.ListPackSizesCalls()))
	})

	t.Run("Large Order With Complex Pack Combination", func(t *testing.T) {
		mockStorage := &mock.PackStorageMock{
			ListPackSizesFunc: func(ctx context.Context) ([]int, error) {
				return []int{23, 31, 53}, nil
			},
		}

		svc := order.NewService(mockStorage)
		orderQty := 500_000
		result, err := svc.PackOrder(t.Context(), orderQty)

		require.NoError(t, err)
		assert.Equal(t, map[int]int{53: 9429, 31: 7, 23: 2}, result)

		total := 0
		for size, qty := range result {
			total += size * qty
		}
		assert.GreaterOrEqual(t, total, orderQty)
		assert.Equal(t, 500_000, total)
	})
}
