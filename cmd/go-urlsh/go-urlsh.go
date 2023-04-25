package app

import (
	"codeberg.org/lauralani/go-urlsh/internal/app"
	"codeberg.org/lauralani/go-urlsh/internal/db"
	"fmt"
	"log"
)

func Run() error {

	err := db.InitializeDB()
	if err != nil {
		log.Fatalln(err)
	}

	err = app.SetupFiber()
	if err != nil {
		return fmt.Errorf("couldn't start webserver: %v", err.Error())
	}
	return nil
}
