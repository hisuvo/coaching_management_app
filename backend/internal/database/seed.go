package database

import (
	"coaching_backend/internal/config"
	"coaching_backend/internal/domain/users"
	"coaching_backend/internal/pkg/security"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB, cnfg *config.Config) error {
	var count int64

	db.Model(&users.User{}).Count(&count)

	if count > 0 {
		return nil
	}

	// hash passwrd generate
	hashPassword, err := security.HashPassword(cnfg.PLATFORM_ADMIN_PASSWORD)

	if err != nil {
		return err
	}

	suerAdmin := users.User{
		Name: cnfg.PLATFORM_ADMIN_NAME,
		Email: cnfg.PLATFORM_ADMIN_EMAIL,
		Password: hashPassword,
		Role: "PLATFORM_ADMIN",
		Phone: cnfg.PLATFORM_ADMIN_PHONE,
	}

	return db.Create(&suerAdmin).Error
}