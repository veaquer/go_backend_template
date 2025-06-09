package migration

import (
	UserModel "backend_template/internal/user/model"
	VerificationModel "backend_template/internal/verification/model"
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
