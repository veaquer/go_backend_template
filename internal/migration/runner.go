package migration

import (
	UserModel "github.com/veaquer/go_backend_template/internal/user/model"
	VerificationModel "github.com/veaquer/go_backend_template/internal/verification/model"
	"log"

	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	err := db.AutoMigrate(
		&UserModel.UserModel{},
		&VerificationModel.EmailVerificationModel{},
		)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
}
