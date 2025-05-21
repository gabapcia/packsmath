package api_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gabapcia/packsmath/internal/handler/api"
	"github.com/gabapcia/packsmath/internal/order/mock"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPackOrderHandler(t *testing.T) {
	t.Run("Returns 200 With Valid Pack Combination", func(t *testing.T) {
		mockService := &mock.ServiceMock{
			PackOrderFunc: func(ctx context.Context, order int) (map[int]int, error) {
				assert.Equal(t, 1250, order)
				return map[int]int{1000: 1, 250: 1}, nil
			},
		}

		app := fiber.New()
		app.Post("/orders", api.PackOrderHandler(mockService))

		body := api.PackOrderRequest{Order: 1250}
		bodyBytes, _ := json.Marshal(body)

		req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewReader(bodyBytes))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var parsed map[int]int
		err = json.NewDecoder(resp.Body).Decode(&parsed)
		require.NoError(t, err)
		assert.Equal(t, map[int]int{1000: 1, 250: 1}, parsed)

		assert.Equal(t, 1, len(mockService.PackOrderCalls()))
	})

	t.Run("Returns 400 On Invalid JSON Body", func(t *testing.T) {
		mockService := &mock.ServiceMock{}

		app := fiber.New(fiber.Config{
			ErrorHandler: api.ErrorHandler,
		})
		app.Post("/orders", api.PackOrderHandler(mockService))

		req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBufferString(`invalid-json`))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		require.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		assert.Empty(t, mockService.PackOrderCalls())
	})

	t.Run("Returns 500 If Service Returns Error", func(t *testing.T) {
		mockService := &mock.ServiceMock{
			PackOrderFunc: func(ctx context.Context, order int) (map[int]int, error) {
				return nil, assert.AnError
			},
		}

		app := fiber.New(fiber.Config{
			ErrorHandler: api.ErrorHandler,
		})
		app.Post("/orders", api.PackOrderHandler(mockService))

		body := api.PackOrderRequest{Order: 100}
		bodyBytes, _ := json.Marshal(body)

		req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewReader(bodyBytes))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		require.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

		assert.Equal(t, 1, len(mockService.PackOrderCalls()))
	})
}
