package main

import (
	"github.com/veaquer/go_backend_template/internal/bootstrap"
	"github.com/veaquer/go_backend_template/internal/migration"
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
