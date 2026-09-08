package students

import "gorm.io/gorm"

type Student struct {
	gorm.Model

	CoachingID string `json:"coaching_id" gorm:"not null;index"`

	Name       string `json:"name" gorm:"type:varchar(100);not null"`
	Class      string `json:"class" gorm:"type:varchar(50);not null"`
	Session    string `json:"session" gorm:"type:varchar(20);not null"`
	Phone      string `json:"phone" gorm:"type:varchar(20);not null"`
	Email      string `json:"email" gorm:"type:varchar(150);not null;uniqueIndex"`
	Address    string `json:"address" gorm:"type:text;not null"`
	SchoolName string `json:"school_name" gorm:"type:varchar(250);not null"`
}