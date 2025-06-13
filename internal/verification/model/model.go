package model

import (
	"time"
)

type EmailVerificationModel struct {
	ID        uint   `gorm:"PrimaryKey"`
	UserID    uint   `gorm:"index"`
	Email     string `gorm:"index"`
	Token     string `gorm:"uniqueIndex"`
	Purpose   string `gorm:"index"` // e.g., "register", "change"
	ExpiresAt time.Time
	CreatedAt time.Time
}
