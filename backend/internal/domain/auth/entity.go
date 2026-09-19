package auth

import (
	"time"

	"gorm.io/gorm"
)

type AuthSession struct {
	gorm.Model

	UserID uint `gorm:"not null;index"`
	CoachingID *uint `gorm:"index"`
	RefreshTokenHash string `gorm:"type:varchar(255);uniqueIndex;not null"`
	UserAgent string `gorm:"type:text"`
	IPAddress string `gorm:"type:varchar(45)"`
	ExpiresAt time.Time `gorm:"not null;index"`
	RevokedAt *time.Time `gorm:"index"`
}