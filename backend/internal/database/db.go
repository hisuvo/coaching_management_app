package database

import (
	"coaching_backend/internal/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB(cnfg *config.Config) *gorm.DB{
	connStr := cnfg.DNS_URL
	logLavel := logger.Error

	if cnfg.APP_ENV == "development" {
		logLavel = logger.Info
	}

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		TranslateError: true,
		Logger: logger.Default.LogMode(logLavel), // authoMigrate info show or not show
	})

	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return nil
	}

	Migrate(db) // it auto create table in db
	fmt.Println("Successfully database connected done!")
	
	return db
}