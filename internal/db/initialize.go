package db

import (
	"codeberg.org/lauralani/go-urlsh/models"
	"context"
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"os"
)

func InitializeDB() error {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv("DATABASE_URL"))))
	models.DB = bun.NewDB(sqldb, pgdialect.New())

	_, err := models.DB.NewCreateTable().IfNotExists().Model((*models.Link)(nil)).Exec(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't create database: [%w]", err)
	}

	_, err = models.DB.NewCreateTable().IfNotExists().Model((*models.User)(nil)).Exec(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't create database: [%w]", err)
	}

	_, err = models.DB.NewCreateTable().IfNotExists().Model((*models.Login)(nil)).Exec(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't create database: [%w]", err)
	}

	_, err = models.DB.NewCreateTable().IfNotExists().Model((*models.ApiKey)(nil)).Exec(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't create database: [%w]", err)
	}
	return nil
}
