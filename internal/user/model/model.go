package model

import (
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	ID              uint   `gorm:"primaryKey"`
	Username        string `gorm:"unique"`
	Name            string `gorm:"non null"`
	Password        string `gorm:"non null"`
	Email           string `gorm:"unique"`
	IsEmailVerified bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
