package main

import (
	app "codeberg.org/lauralani/go-urlsh/cmd/go-urlsh"
	"log"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
