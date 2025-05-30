// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"github.com/gabapcia/packsmath/internal/order"
	"sync"
)

// Ensure, that ServiceMock does implement order.Service.
// If this is not the case, regenerate this file with moq.
var _ order.Service = &ServiceMock{}

// ServiceMock is a mock implementation of order.Service.
//
//	func TestSomethingThatUsesService(t *testing.T) {
//
//		// make and configure a mocked order.Service
//		mockedService := &ServiceMock{
//			PackOrderFunc: func(ctx context.Context, order int) (map[int]int, error) {
//				panic("mock out the PackOrder method")
//			},
//		}
//
//		// use mockedService in code that requires order.Service
//		// and then make assertions.
//
//	}
type ServiceMock struct {
	// PackOrderFunc mocks the PackOrder method.
	PackOrderFunc func(ctx context.Context, order int) (map[int]int, error)

	// calls tracks calls to the methods.
	calls struct {
		// PackOrder holds details about calls to the PackOrder method.
		PackOrder []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Order is the order argument value.
			Order int
		}
	}
	lockPackOrder sync.RWMutex
}

// PackOrder calls PackOrderFunc.
func (mock *ServiceMock) PackOrder(ctx context.Context, order int) (map[int]int, error) {
	if mock.PackOrderFunc == nil {
		panic("ServiceMock.PackOrderFunc: method is nil but Service.PackOrder was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Order int
	}{
		Ctx:   ctx,
		Order: order,
	}
	mock.lockPackOrder.Lock()
	mock.calls.PackOrder = append(mock.calls.PackOrder, callInfo)
	mock.lockPackOrder.Unlock()
	return mock.PackOrderFunc(ctx, order)
}

// PackOrderCalls gets all the calls that were made to PackOrder.
// Check the length with:
//
//	len(mockedService.PackOrderCalls())
func (mock *ServiceMock) PackOrderCalls() []struct {
	Ctx   context.Context
	Order int
} {
	var calls []struct {
		Ctx   context.Context
		Order int
	}
	mock.lockPackOrder.RLock()
	calls = mock.calls.PackOrder
	mock.lockPackOrder.RUnlock()
	return calls
}
