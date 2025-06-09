package main

import (
	"backend_template/internal/bootstrap"
	"log"
)

func main() {
	app, err := bootstrap.Bootstrap()
	if err != nil {
		log.Fatalf("Bootstrap error: %v", err)
	}

	app.Router.Run()
}
