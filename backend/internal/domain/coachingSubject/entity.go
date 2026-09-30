package coachingsubject

import (
	"gorm.io/gorm"
)

type CoachingSubjectStatus string

const (
	CoachingSubjectStatusActive   CoachingSubjectStatus = "ACTIVE"
	CoachingSubjectStatusInactive CoachingSubjectStatus = "INACTIVE"
)

type CoachingSubject struct {
	gorm.Model

	CoachingID uint `gorm:"not null;index"`
	SubjectID  uint `gorm:"not null;index"`

	Status string `gorm:"type:varchar(20);not null;default:'ACTIVE';index"`
}
