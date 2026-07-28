package links

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/lila-network/go-urlsh/internal/fault"
	"github.com/lila-network/go-urlsh/internal/misc"
	"github.com/uptrace/bun"
)

// Repository provides access to link-related database operations.
type Repository interface {
	GetAll(ctx context.Context) ([]Link, error)
	GetByID(ctx context.Context, id string) (Link, error)
	Create(ctx context.Context, newLink Link) (Link, error)
	Update(ctx context.Context, link Link) (Link, error)
	DeleteByID(ctx context.Context, id string) error
}

// LinkRepository provides access to link-related database operations, such as retrieving, creating, updating, and deleting links.
type LinkRepository struct {
	database *bun.DB
}

// NewLinkRepository creates a new instance of LinkRepository with the provided database connection.
func NewLinkRepository(database *bun.DB) Repository {
	return &LinkRepository{
		database: database,
	}
}

// GetAll retrieves all links from the database using the given context. Returns the list of links or an error if the operation fails.
func (l *LinkRepository) GetAll(ctx context.Context) (links []Link, err error) {
	links = []Link{}

	err = l.database.NewSelect().Model(&links).Scan(ctx)
	if err != nil {
		log.Printf("[HandleLinkGetAll] Failed to get links from DB: %q", err)
		return
	}

	return
}

// GetByID retrieves a link from the database using the given ID and context. Returns the link or an error if not found.
func (l *LinkRepository) GetByID(ctx context.Context, id string) (link Link, err error) {
	if id == "" {
		err = fault.ErrInvalidParameters
		return
	}

	err = l.database.NewSelect().Model(&link).Where("id = ?", id).Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Handle the "not found" case specifically
			return link, fault.ErrNotFound
		}
		// Handle other database errors (connection issues, etc.)
		return link, err
	}

	return
}

// Create creates a new link in the database using the given context and link. Returns the created link or an error if the link already exists.
func (l *LinkRepository) Create(ctx context.Context, newLink Link) (link Link, err error) {
	const maxAttempts = 10

	customSlug := newLink.ID != ""

	now := time.Now().UTC()
	newLink.Created = now
	newLink.Modified = now

	for range maxAttempts {
		if !customSlug {
			newLink.ID = misc.RandomString(6)
		}

		result, err := l.database.
			NewInsert().
			Model(&newLink).
			Ignore().
			Exec(ctx)
		if err != nil {
			return Link{}, fmt.Errorf(
				"failed to insert link %q: %w",
				newLink.ID,
				err,
			)
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return Link{}, fmt.Errorf(
				"get affected rows after inserting link %q: %w",
				newLink.ID,
				err,
			)
		}

		if rowsAffected > 0 {
			return newLink, nil
		}

		if customSlug {
			return Link{}, fault.ErrAlreadyExists
		}
	}

	return Link{}, fmt.Errorf(
		"failed to generate unique link ID after %d attempts",
		maxAttempts,
	)
}

// Update modifies an existing link in the database and returns the updated link or an error if the operation fails.
func (l *LinkRepository) Update(ctx context.Context, link Link) (Link, error) {
	link.Modified = time.Now().UTC()

	result, err := l.database.NewUpdate().
		Model(&link).
		WherePK().
		Exec(ctx)
	if err != nil {
		return link, fmt.Errorf("failed to update link: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return link, fmt.Errorf("failed to retrieve rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return link, fault.ErrNotFound
	}

	return link, nil
}

// DeleteByID deletes an existing link from the database using the given context and link. Returns an error if the link does not exist.
func (l *LinkRepository) DeleteByID(ctx context.Context, id string) error {
	if id == "" {
		return fault.ErrInvalidParameters
	}

	link := Link{
		ID: id,
	}

	result, err := l.database.NewDelete().Model(&link).WherePK().Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete link: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to delete link: %w", err)
	}

	if rowsAffected == 0 {
		return fault.ErrNotFound
	}

	return nil
}
