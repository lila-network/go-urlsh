package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/lila-network/go-urlsh/internal/links"
	"github.com/lila-network/go-urlsh/models"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

// InitializeDB initializes the database connection, sets up tables, and applies schema migrations for the application.
func InitializeDB() (*bun.DB, error) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv("DATABASE_URL"))))
	models.DB = bun.NewDB(sqldb, pgdialect.New())

	_, err := models.DB.NewCreateTable().IfNotExists().Model((*links.Link)(nil)).Exec(context.Background())
	if err != nil {
		return models.DB, fmt.Errorf("couldn't create database: [%w]", err)
	}

	_, err = models.DB.NewCreateTable().IfNotExists().Model((*models.User)(nil)).Exec(context.Background())
	if err != nil {
		return models.DB, fmt.Errorf("couldn't create database: [%w]", err)
	}

	_, err = models.DB.NewCreateTable().IfNotExists().Model((*models.Session)(nil)).Exec(context.Background())
	if err != nil {
		return models.DB, fmt.Errorf("couldn't create database: [%w]", err)
	}

	_, err = models.DB.NewCreateTable().IfNotExists().Model((*models.ApiKey)(nil)).Exec(context.Background())
	if err != nil {
		return models.DB, fmt.Errorf("couldn't create database: [%w]", err)
	}
	return models.DB, nil
}
