package main

import (
	"log"
	"os"

	"github.com/re-partners-challenge-backend/app"
)

func main() {

	app, cleanup, err := app.Build()
	if err != nil {
		log.Fatalf("Error on build project. Error: %s", err.Error())
	}

	if err := app.Server().Start(); err != nil {
		cleanup()
		log.Fatalf("failed to start the http server. Error: %s", err.Error())
	}

	cleanup()
	os.Exit(0)
}
