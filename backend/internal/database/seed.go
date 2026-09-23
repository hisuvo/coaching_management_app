package database

import (
	"coaching_backend/internal/config"
	"coaching_backend/internal/domain/auth"
	"coaching_backend/internal/domain/users"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB, cnfg *config.Config) error {
	var count int64

	db.Model(&users.User{}).Count(&count)

	if count > 0 {
		return nil
	}

	// hash passwrd generate
	hashPassword, err := auth.HashPassword(cnfg.SUPER_ADMIN_PASSWORD)

	if err != nil {
		return err
	}

	suerAdmin := users.User{
		Name: cnfg.SUPER_ADMIN_NAME,
		Email: cnfg.SUPER_ADMIN_EMAIL,
		PasswordHash: hashPassword,
		Role: "SUPER_ADMIN",
		Phone: cnfg.SUPER_ADMIN_PHONE,
	}

	return db.Create(&suerAdmin).Error
}