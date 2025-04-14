package main

import (
	"fmt"
	"log"

	"codeberg.org/lauralani/go-urlsh/internal/app"
	"codeberg.org/lauralani/go-urlsh/internal/db"
)

func main() {
	err := Run()
	if err != nil {
		log.Fatal(err)
	}
}

func Run() error {

	err := db.InitializeDB()
	if err != nil {
		log.Fatalln(err)
	}

	go app.CleanupLogins()
	go app.CleanupLoginsCronJob()

	err = app.SetupFiber()
	if err != nil {
		return fmt.Errorf("couldn't start webserver: %v", err.Error())
	}
	return nil
}
