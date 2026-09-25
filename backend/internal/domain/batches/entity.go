package batches

import (
	"time"

	"gorm.io/gorm"
)

type BranchStatus string

var (
	BranchStatusActive   BranchStatus = "ACTIVE"
	BranchStatusInActive BranchStatus = "INACTIVE"
)

type Branch struct {
	gorm.Model

	// Relationship with the parent coaching.
	CoachingID uint `gorm:"not null;index;uniqueIndex:idx_coaching_branch_code"`

	// Basic branch information.
	Name    string `gorm:"type:varchar(150);not null"`
	Code    string `gorm:"type:varchar(50);not null;uniqueIndex:idx_coaching_branch_code"`
	Phone   string `gorm:"type:varchar(20)"`
	Email   string `gorm:"type:varchar(150)"`
	Address string `gorm:"type:text"`

	// Branch location.
	City    string `gorm:"type:varchar(100)"`
	Division string `gorm:"type:varchar(100)"`

	// Branch lifecycle status.
	Status BranchStatus `gorm:"type:varchar(20);not null;default:'ACTIVE';index"`

	// Optional operating timezone.
	TimeZone string `gorm:"type:varchar(100);not null;default:'Asia/Dhaka'"`

	// Optional opening/closing information.	
	OpeningTime *time.Time `gorm:"type:time"`
	ClosingTime *time.Time `gorm:"type:time"`
}

//* Note:
// Entity → GORM tags
// Request DTO → JSON + validation tags
// Response DTO → JSON tags