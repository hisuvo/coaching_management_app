package database

import (
	"coaching_backend/internal/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB(cnfg *config.Config) *gorm.DB{
	// connStr := cnfg.DNS_URL
	logLavel := logger.Error

	if cnfg.APP_ENV == "development" {
		logLavel = logger.Info
	}

	dns := "host=localhost user=postgres password=12345 dbname=coaching_management port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		TranslateError: true,
		Logger: logger.Default.LogMode(logLavel), // authoMigrate info show or not show
	})

	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return nil
	}

	// Migrate auto create table in database
	Migrate(db) 

	// Super Admin seed function call here
	if err := Seed(db, cnfg); err == nil {
		fmt.Println("Super Admin seed completed successfully")
	}else{
		fmt.Println("Super Admin seed failed")
	}

	fmt.Println("Successfully database connected done!")
	return db
}