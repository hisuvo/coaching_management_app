package assignments

import (
	"time"

	"gorm.io/gorm"
)

type AssignmentStatus string

const (
	AssignmentStatusDraft     AssignmentStatus = "DRAFT"
	AssignmentStatusPublished AssignmentStatus = "PUBLISHED"
	AssignmentStatusClosed    AssignmentStatus = "CLOSED"
)

type Assignment struct {
	gorm.Model

	Title       string           `gorm:"type:varchar(200);not null"`
	Description *string          `gorm:"type:text"`

	// here use indexs beacuse can frequently query
	SubjectID   uint             `gorm:"not null;index"`
	CoachingID  uint             `gorm:"not null;index"`
	CreatedBy   uint             `gorm:"not null;index"` // Don't accept CreatedBy from the client
	DueDate     *time.Time       `gorm:"index"`
	Status      AssignmentStatus `gorm:"type:varchar(20);not null;default:'draft';index"`
}