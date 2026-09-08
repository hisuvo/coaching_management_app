package coaching

import (
	"coaching_backend/internal/domain/students"
	"coaching_backend/internal/domain/teachers"

	"gorm.io/gorm"
)

type Coaching struct {
	gorm.Model

	Name   string `json:"name" gorm:"type:varchar(150);not null"`
	Email  string `json:"email" gorm:"type:varchar(150);unique;not null"`
	Phone  string `json:"phone" gorm:"type:varchar(20);not null"`
	Domain string `json:"domain" gorm:"type:varchar(255);unique;not null"`

	Teachers []teachers.Teacher `json:"teachers,omitempty"`
	Students []students.Student `json:"students,omitempty"`
}