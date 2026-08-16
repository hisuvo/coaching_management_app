package database

import (
	"fmt"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	fmt.Println("Migrate info ->", db)
}