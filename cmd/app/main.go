package main

import (
	"log"

	"todo/actions"
	"todo/models"
)

func main() {
	if err := models.Migrate(); err != nil {
		log.Fatal(err)
	}

	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
