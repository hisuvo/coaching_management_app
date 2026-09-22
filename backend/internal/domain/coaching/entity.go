package coaching

import (
	"gorm.io/gorm"
)

type CoachingStatus string

const (
	StatusActive   CoachingStatus = "active"
	StatusInactive CoachingStatus = "inactive"
	StatusSuspended CoachingStatus = "suspended"
)

type Coaching struct {
	gorm.Model

	Name     string `gorm:"type:varchar(150);not null"`
	Email    string `gorm:"type:varchar(150);uniqueIndex;not null"`
	Phone    string `gorm:"type:varchar(20);not null"`
	Domain   string `gorm:"type:varchar(255);uniqueIndex;not null"`
	Slug     string `gorm:"type:varchar(100);uniqueIndex;not null"`
	LogoURL  string `gorm:"type:text"`
	Address  string `gorm:"type:text"`
	TimeZone string `gorm:"type:varchar(50);not null;default:'Asia/Dhaka'"`
	Status   CoachingStatus `gorm:"type:varchar(20);not null;default:'active';check:status IN ('active','inactive','suspended')"`
}