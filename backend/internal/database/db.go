package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB{
	dns := "host=localhost user=postgres password=12345 dbname=coaching_management port=5000 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return nil
	}

	Migrate(db)
	fmt.Println("Successfully database connected done!")
	return db
}