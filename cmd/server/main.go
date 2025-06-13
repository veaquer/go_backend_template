package main

import (
	"backend_template/internal/bootstrap"
	"backend_template/internal/migration"
	"log"
)

func main() {
	app, err := bootstrap.Bootstrap()
	if err != nil {
		log.Fatalf("Bootstrap error: %v", err)
	}
	migration.Run(app.DB)
	app.Router.Run()
}
