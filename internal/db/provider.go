package db

import (
	"github.com/veaquer/go_backend_template/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ProvideDB(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
