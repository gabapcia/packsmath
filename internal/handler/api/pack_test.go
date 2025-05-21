package api_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gabapcia/packsmath/internal/handler/api"
	"github.com/gabapcia/packsmath/internal/pack"
	"github.com/gabapcia/packsmath/internal/pack/mock"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegisterPackSizeHandler(t *testing.T) {
	t.Run("Returns 204 On Success", func(t *testing.T) {
		mockService := &mock.ServiceMock{
			RegisterPackSizeFunc: func(ctx context.Context, size int) error {
				assert.Equal(t, 500, size)
				return nil
			},
		}

		app := fiber.New()
		app.Post("/packs", api.RegisterPackSizeHandler(mockService))

		body := api.RegisterPackRequest{Size: 500}
		bodyBytes, _ := json.Marshal(body)

		req := httptest.NewRequest(http.MethodPost, "/packs", bytes.NewReader(bodyBytes))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		require.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, resp.StatusCode)

		assert.Equal(t, 1, len(mockService.RegisterPackSizeCalls()))
	})

	t.Run("Returns 400 On Invalid JSON Body", func(t *testing.T) {
		mockService := &mock.ServiceMock{}

		app := fiber.New(fiber.Config{
			ErrorHandler: api.ErrorHandler,
		})
		app.Post("/packs", api.RegisterPackSizeHandler(mockService))

		req := httptest.NewRequest(http.MethodPost, "/packs", bytes.NewBufferString(`invalid-json`))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		require.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		assert.Empty(t, mockService.RegisterPackSizeCalls())
	})

	t.Run("Returns 409 When Pack Already Registered", func(t *testing.T) {
		mockService := &mock.ServiceMock{
			RegisterPackSizeFunc: func(ctx context.Context, size int) error {
				return pack.ErrPackSizeAlreadyRegistered
			},
		}

		app := fiber.New(fiber.Config{
			ErrorHandler: api.ErrorHandler,
		})
		app.Post("/packs", api.RegisterPackSizeHandler(mockService))

		body := api.RegisterPackRequest{Size: 250}
		bodyBytes, _ := json.Marshal(body)

		req := httptest.NewRequest(http.MethodPost, "/packs", bytes.NewReader(bodyBytes))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		require.NoError(t, err)
		assert.Equal(t, http.StatusConflict, resp.StatusCode)

		assert.Equal(t, 1, len(mockService.RegisterPackSizeCalls()))
	})

	t.Run("Returns 500 On Unknown Error", func(t *testing.T) {
		mockService := &mock.ServiceMock{
			RegisterPackSizeFunc: func(ctx context.Context, size int) error {
				return errors.New("unexpected error")
			},
		}

		app := fiber.New(fiber.Config{
			ErrorHandler: api.ErrorHandler,
		})
		app.Post("/packs", api.RegisterPackSizeHandler(mockService))

		body := api.RegisterPackRequest{Size: 999}
		bodyBytes, _ := json.Marshal(body)

		req := httptest.NewRequest(http.MethodPost, "/packs", bytes.NewReader(bodyBytes))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		require.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

		assert.Equal(t, 1, len(mockService.RegisterPackSizeCalls()))
	})
}

func TestListPackSizesHandler(t *testing.T) {
	t.Run("Returns 200 With List Of Pack Sizes", func(t *testing.T) {
		mockService := &mock.ServiceMock{
			ListPackSizesFunc: func(ctx context.Context) ([]int, error) {
				return []int{250, 500, 1000}, nil
			},
		}

		app := fiber.New()
		app.Get("/packs", api.ListPackSizesHandler(mockService))

		req := httptest.NewRequest(http.MethodGet, "/packs", nil)

		resp, err := app.Test(req)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var result []int
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)
		assert.Equal(t, []int{250, 500, 1000}, result)

		assert.Equal(t, 1, len(mockService.ListPackSizesCalls()))
	})

	t.Run("returns 500 on service error", func(t *testing.T) {
		mockService := &mock.ServiceMock{
			ListPackSizesFunc: func(ctx context.Context) ([]int, error) {
				return nil, errors.New("storage failure")
			},
		}

		app := fiber.New(fiber.Config{
			ErrorHandler: api.ErrorHandler,
		})
		app.Get("/packs", api.ListPackSizesHandler(mockService))

		req := httptest.NewRequest(http.MethodGet, "/packs", nil)

		resp, err := app.Test(req)
		require.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

		assert.Equal(t, 1, len(mockService.ListPackSizesCalls()))
	})
}

func TestDeletePackSizeHandler(t *testing.T) {
	t.Run("Returns 204 On Successful Delete", func(t *testing.T) {
		mockService := &mock.ServiceMock{
			DeletePackSizeFunc: func(ctx context.Context, size int) error {
				assert.Equal(t, 500, size)
				return nil
			},
		}

		app := fiber.New()
		app.Delete("/packs/:size", api.DeletePackSizeHandler(mockService))

		req := httptest.NewRequest(http.MethodDelete, "/packs/500", nil)

		resp, err := app.Test(req)
		require.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, resp.StatusCode)
		assert.Equal(t, 1, len(mockService.DeletePackSizeCalls()))
	})

	t.Run("Returns 404 When Pack Not Found", func(t *testing.T) {
		mockService := &mock.ServiceMock{
			DeletePackSizeFunc: func(ctx context.Context, size int) error {
				return pack.ErrPackSizeNotFound
			},
		}

		app := fiber.New(fiber.Config{
			ErrorHandler: api.ErrorHandler,
		})
		app.Delete("/packs/:size", api.DeletePackSizeHandler(mockService))

		req := httptest.NewRequest(http.MethodDelete, "/packs/999", nil)

		resp, err := app.Test(req)
		require.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
		assert.Equal(t, 1, len(mockService.DeletePackSizeCalls()))
	})

	t.Run("returns 500 on unexpected error", func(t *testing.T) {
		mockService := &mock.ServiceMock{
			DeletePackSizeFunc: func(ctx context.Context, size int) error {
				return errors.New("something went wrong")
			},
		}

		app := fiber.New(fiber.Config{
			ErrorHandler: api.ErrorHandler,
		})
		app.Delete("/packs/:size", api.DeletePackSizeHandler(mockService))

		req := httptest.NewRequest(http.MethodDelete, "/packs/123", nil)

		resp, err := app.Test(req)
		require.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		assert.Equal(t, 1, len(mockService.DeletePackSizeCalls()))
	})
}
