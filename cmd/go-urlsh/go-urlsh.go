package app

import (
	"codeberg.org/lauralani/go-urlsh/internal/app"
	"codeberg.org/lauralani/go-urlsh/models"
	"database/sql"
	"fmt"
	"os"
)

func Run() error {
	var dberr error
	models.DB, dberr = sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if dberr != nil {
		return fmt.Errorf("couldn't connect to database: %v", dberr.Error())
	}

	err := app.SetupFiber()
	if err != nil {
		return fmt.Errorf("couldn't start webserver: %v", err.Error())
	}
	return nil
}
