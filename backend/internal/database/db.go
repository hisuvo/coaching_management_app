package database

import (
	"coaching_backend/internal/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cnfg *config.Config) *gorm.DB{
	connStr := cnfg.DNS_URL

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return nil
	}

	Migrate(db) // it auto create table in db
	fmt.Println("Successfully database connected done!")
	
	return db
}