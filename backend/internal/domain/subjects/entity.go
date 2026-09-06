package subjects

import "gorm.io/gorm"

type SubjectStatus string

const (
	SubjectStatusActive   SubjectStatus = "active"
	SubjectStatusInactive SubjectStatus = "inactive"
)

type Subject struct {
	gorm.Model

	Name string `gorm:"type:varchar(100);not null" json:"name"`
	Code string `gorm:"type:varchar(50);not null" json:"code"`
	Description *string `gorm:"type:text" json:"description,omitempty"`
	Status SubjectStatus `gorm:"type:varchar(20);not null;default:'active'" json:"status"`
}