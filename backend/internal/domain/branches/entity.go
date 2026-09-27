package branches

import (
	"gorm.io/gorm"
)

type BranchStatus string

const (
	BranchStatusActive   BranchStatus = "ACTIVE"
	BranchStatusInactive BranchStatus = "INACTIVE"
)

type Branch struct {
	gorm.Model

	CoachingID uint `gorm:"not null;index"`

	Name string `gorm:"type:varchar(100);not null"`
	Code string `gorm:"type:varchar(50);uniqueIndex;not null"`
	Phone string `gorm:"type:varchar(20)"`
	Email string `gorm:"type:varchar(150);email"`
	Address string `gorm:"type:text"`

	Status BranchStatus `gorm:"type:varchar(20);default:'ACTIVE';index"`
}