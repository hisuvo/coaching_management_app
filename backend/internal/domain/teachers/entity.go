package teachers

import (
	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model

	CoachingID uint `json:"coaching_id" gorm:"not null;index"`

	Name        string `json:"name" gorm:"type:varchar(100);not null"`
	Email       string `json:"email" gorm:"type:varchar(150);not null"`
	Phone       string `json:"phone" gorm:"type:varchar(20);not null"`
	Password    string `json:"-" gorm:"type:varchar(255);not null"`
	Subject     string `json:"subject" gorm:"type:varchar(100);not null"`
	Designation string `json:"designation" gorm:"type:varchar(100)"`
	Address     string `json:"address" gorm:"type:text"`
	IsActive    bool   `json:"is_active" gorm:"default:true"`
}

/*
{
  "coaching_id": 1,
  "name": "Rahim Ahmed",
  "email": "rahim.ahmed@example.com",
  "phone": "01712345678",
  "password": "12345678",
  "subject": "Mathematics",
  "designation": "Senior Teacher",
  "address": "Dhaka, Bangladesh",
  "is_active": true
}
*/