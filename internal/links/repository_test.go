package links

import (
	"context"
	"testing"

	"github.com/uptrace/bun"
)

func TestNewLinkRepository(t *testing.T) {
	db := &bun.DB{}
	repo := NewLinkRepository(db)
	if repo == nil {
		t.Fatal("NewLinkRepository returned nil")
	}

	_, ok := repo.(Repository)
	if !ok {
		t.Error("NewLinkRepository() does not return a Repository interface")
	}
}

// MockRepository is a mock implementation of the Repository interface.
type MockRepository struct {
	GetAllFn     func(ctx context.Context) ([]Link, error)
	GetByIDFn    func(ctx context.Context, id string) (Link, error)
	CreateFn     func(ctx context.Context, newLink Link) (Link, error)
	UpdateFn     func(ctx context.Context, link Link) (Link, error)
	DeleteByIDFn func(ctx context.Context, id string) error
}

func (m *MockRepository) GetAll(ctx context.Context) ([]Link, error) {
	return m.GetAllFn(ctx)
}

func (m *MockRepository) GetByID(ctx context.Context, id string) (Link, error) {
	return m.GetByIDFn(ctx, id)
}

func (m *MockRepository) Create(ctx context.Context, newLink Link) (Link, error) {
	return m.CreateFn(ctx, newLink)
}

func (m *MockRepository) Update(ctx context.Context, link Link) (Link, error) {
	return m.UpdateFn(ctx, link)
}

func (m *MockRepository) DeleteByID(ctx context.Context, id string) error {
	return m.DeleteByIDFn(ctx, id)
}
