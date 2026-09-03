package subjects

import (
	"gorm.io/gorm"
)

type SubjectStatus string

const (
	SubjectStatusActive   SubjectStatus = "active"
	SubjectStatusInactive SubjectStatus = "inactive"
)

type Subject struct {
	gorm.Model
	Name        string `gorm:"type:vacher(100);not null" json:"name"`
	Code        string `gorm:"type:vacher(50);not null" json:"code"`
	Description *string `gorm:"type:text" json:"description;omitempty"`
	Status      SubjectStatus `gorm:"type:vacher(20);not null;degault:'active'" json:"status"`
}