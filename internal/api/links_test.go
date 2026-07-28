package api

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/lila-network/go-urlsh/internal/fault"
	"github.com/lila-network/go-urlsh/internal/links"
)

// mockRepository is a mock implementation of the Repository interface for testing.
type mockRepository struct {
	GetAllFn     func(ctx context.Context) ([]links.Link, error)
	GetByIDFn    func(ctx context.Context, id string) (links.Link, error)
	CreateFn     func(ctx context.Context, newLink links.Link) (links.Link, error)
	UpdateFn     func(ctx context.Context, link links.Link) (links.Link, error)
	DeleteByIDFn func(ctx context.Context, id string) error
}

func (m *mockRepository) GetAll(ctx context.Context) ([]links.Link, error) {
	return m.GetAllFn(ctx)
}

func (m *mockRepository) GetByID(ctx context.Context, id string) (links.Link, error) {
	return m.GetByIDFn(ctx, id)
}

func (m *mockRepository) Create(ctx context.Context, newLink links.Link) (links.Link, error) {
	return m.CreateFn(ctx, newLink)
}

func (m *mockRepository) Update(ctx context.Context, link links.Link) (links.Link, error) {
	return m.UpdateFn(ctx, link)
}

func (m *mockRepository) DeleteByID(ctx context.Context, id string) error {
	return m.DeleteByIDFn(ctx, id)
}

func TestLinkHandler_GetAll(t *testing.T) {
	app := fiber.New()
	mockRepo := &mockRepository{}
	handler := NewLinkHandler(mockRepo)
	app.Get("/links", handler.GetAll)

	t.Run("success", func(t *testing.T) {
		expectedLinks := []links.Link{
			{ID: "test1", URL: "https://example.com/1"},
			{ID: "test2", URL: "https://example.com/2"},
		}
		mockRepo.GetAllFn = func(ctx context.Context) ([]links.Link, error) {
			return expectedLinks, nil
		}

		req := httptest.NewRequest(http.MethodGet, "/links", nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("app.Test failed: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("status = %v, want %v", resp.StatusCode, http.StatusOK)
		}

		var body map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&body)
		if body["status"] != "success" {
			t.Errorf("body[status] = %v, want success", body["status"])
		}
	})

	t.Run("not found", func(t *testing.T) {
		mockRepo.GetAllFn = func(ctx context.Context) ([]links.Link, error) {
			return nil, fault.ErrNotFound
		}

		req := httptest.NewRequest(http.MethodGet, "/links", nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("app.Test failed: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("status = %v, want %v", resp.StatusCode, http.StatusOK)
		}
	})

	t.Run("internal error", func(t *testing.T) {
		mockRepo.GetAllFn = func(ctx context.Context) ([]links.Link, error) {
			return nil, errors.New("db error")
		}

		req := httptest.NewRequest(http.MethodGet, "/links", nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("app.Test failed: %v", err)
		}

		if resp.StatusCode != http.StatusInternalServerError {
			t.Errorf("status = %v, want %v", resp.StatusCode, http.StatusInternalServerError)
		}
	})
}

func TestLinkHandler_Get(t *testing.T) {
	app := fiber.New()
	mockRepo := &mockRepository{}
	handler := NewLinkHandler(mockRepo)
	app.Get("/links/:id", handler.Get)

	t.Run("success", func(t *testing.T) {
		expectedLink := links.Link{ID: "test1", URL: "https://example.com/1"}
		mockRepo.GetByIDFn = func(ctx context.Context, id string) (links.Link, error) {
			if id != "test1" {
				return links.Link{}, fault.ErrNotFound
			}
			return expectedLink, nil
		}

		req := httptest.NewRequest(http.MethodGet, "/links/test1", nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("app.Test failed: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("status = %v, want %v", resp.StatusCode, http.StatusOK)
		}
	})

	t.Run("not found", func(t *testing.T) {
		mockRepo.GetByIDFn = func(ctx context.Context, id string) (links.Link, error) {
			return links.Link{}, fault.ErrNotFound
		}

		req := httptest.NewRequest(http.MethodGet, "/links/unknown", nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("app.Test failed: %v", err)
		}

		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("status = %v, want %v", resp.StatusCode, http.StatusNotFound)
		}
	})
}

func TestLinkHandler_Delete(t *testing.T) {
	app := fiber.New()
	mockRepo := &mockRepository{}
	handler := NewLinkHandler(mockRepo)
	app.Delete("/links/:id", handler.Delete)

	t.Run("success", func(t *testing.T) {
		mockRepo.DeleteByIDFn = func(ctx context.Context, id string) error {
			return nil
		}

		req := httptest.NewRequest(http.MethodDelete, "/links/test1", nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("app.Test failed: %v", err)
		}

		if resp.StatusCode != http.StatusNoContent {
			t.Errorf("status = %v, want %v", resp.StatusCode, http.StatusNoContent)
		}
	})

	t.Run("not found", func(t *testing.T) {
		mockRepo.DeleteByIDFn = func(ctx context.Context, id string) error {
			return fault.ErrNotFound
		}

		req := httptest.NewRequest(http.MethodDelete, "/links/unknown", nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("app.Test failed: %v", err)
		}

		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("status = %v, want %v", resp.StatusCode, http.StatusNotFound)
		}
	})
}

func TestLinkHandler_Post(t *testing.T) {
	app := fiber.New()
	mockRepo := &mockRepository{}
	handler := NewLinkHandler(mockRepo)
	app.Post("/links", handler.Post)

	t.Run("success", func(t *testing.T) {
		mockRepo.CreateFn = func(ctx context.Context, newLink links.Link) (links.Link, error) {
			return newLink, nil
		}

		body := `{"id": "new-link", "url": "https://example.com"}`
		req := httptest.NewRequest(http.MethodPost, "/links", ioReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("app.Test failed: %v", err)
		}

		if resp.StatusCode != http.StatusCreated {
			t.Errorf("status = %v, want %v", resp.StatusCode, http.StatusCreated)
		}
	})

	t.Run("invalid request - missing URL", func(t *testing.T) {
		body := `{"id": "new-link"}`
		req := httptest.NewRequest(http.MethodPost, "/links", ioReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("app.Test failed: %v", err)
		}

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("status = %v, want %v", resp.StatusCode, http.StatusBadRequest)
		}
	})

	t.Run("conflict", func(t *testing.T) {
		mockRepo.CreateFn = func(ctx context.Context, newLink links.Link) (links.Link, error) {
			return links.Link{}, fault.ErrAlreadyExists
		}

		body := `{"id": "existing", "url": "https://example.com"}`
		req := httptest.NewRequest(http.MethodPost, "/links", ioReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("app.Test failed: %v", err)
		}

		if resp.StatusCode != http.StatusConflict {
			t.Errorf("status = %v, want %v", resp.StatusCode, http.StatusConflict)
		}
	})
}

func TestLinkHandler_Patch(t *testing.T) {
	app := fiber.New()
	mockRepo := &mockRepository{}
	handler := NewLinkHandler(mockRepo)
	app.Patch("/links/:id", handler.Patch)

	t.Run("success - full update", func(t *testing.T) {
		existingLink := links.Link{ID: "test1", URL: "https://old.com"}
		mockRepo.GetByIDFn = func(ctx context.Context, id string) (links.Link, error) {
			return existingLink, nil
		}
		mockRepo.UpdateFn = func(ctx context.Context, link links.Link) (links.Link, error) {
			return link, nil
		}

		body := `{"url": "https://new.com", "description": "new desc"}`
		req := httptest.NewRequest(http.MethodPatch, "/links/test1", ioReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("app.Test failed: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("status = %v, want %v", resp.StatusCode, http.StatusOK)
		}
	})

	t.Run("not found", func(t *testing.T) {
		mockRepo.GetByIDFn = func(ctx context.Context, id string) (links.Link, error) {
			return links.Link{}, fault.ErrNotFound
		}

		body := `{"url": "https://new.com"}`
		req := httptest.NewRequest(http.MethodPatch, "/links/unknown", ioReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("app.Test failed: %v", err)
		}

		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("status = %v, want %v", resp.StatusCode, http.StatusNotFound)
		}
	})
}

func ioReader(s string) io.Reader {
	return strings.NewReader(s)
}
