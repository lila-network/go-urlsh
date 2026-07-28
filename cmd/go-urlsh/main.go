package main

import (
	"fmt"
	"log"

	"github.com/lila-network/go-urlsh/internal/app"
	"github.com/lila-network/go-urlsh/internal/db"
)

func main() {
	err := Run()
	if err != nil {
		log.Fatal(err)
	}
}

func Run() error {

	database, err := db.InitializeDB()
	if err != nil {
		log.Fatalln(err)
	}

	go app.CleanupLogins()
	go app.CleanupLoginsCronJob()

	err = app.SetupFiber(database)
	if err != nil {
		return fmt.Errorf("couldn't start webserver: %v", err.Error())
	}
	return nil
}
