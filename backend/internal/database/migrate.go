package database

import (
	"coaching_backend/internal/domain/users"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&users.User{},
	)
}