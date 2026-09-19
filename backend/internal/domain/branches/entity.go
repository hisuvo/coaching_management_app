package branches

import (
	"gorm.io/gorm"
)

type Branch struct {
	gorm.Model

	CoachingID uint `json:"coaching_id" gorm:"not null;index"`

	Name string `json:"name" gorm:"type:varchar(100);not null"`
	Code string `json:"code" gorm:"type:varchar(50);not null"`
	Phone string `json:"phone" gorm:"type:varchar(20)"`
	Email string `json:"email" gorm:"type:varchar(150)"`
	Address string `json:"address" gorm:"type:text"`
	City string `json:"city" gorm:"type:varchar(100)"`
	Country string `json:"country" gorm:"type:varchar(100);default:'Bangladesh'"`
	Status string `json:"status" gorm:"type:varchar(20);default:'active';index"`

}