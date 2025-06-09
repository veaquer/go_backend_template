package model

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Username        string `gorm:"unique"`
	Name            string `gorm:"non null"`
	Password        string `gorm:"non null"`
	Email           string `gorm:"unique"`
	IsEmailVerified bool
}
