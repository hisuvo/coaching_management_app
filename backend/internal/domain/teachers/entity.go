package teachers

import (
	"coaching_backend/internal/domain/branches"
	"coaching_backend/internal/domain/coaching"
	"coaching_backend/internal/domain/subjects"
	"coaching_backend/internal/domain/users"
	"time"

	"gorm.io/gorm"
)

type TeacherStatus string

const (
	TeacherStatusActive TeacherStatus = "ACTIVE"
	TeacherStatusInactive TeacherStatus = "INACTIVE"
)

type Teacher struct {
	gorm.Model

	CoachingID uint `gorm:"not null;index"`
	UserID     uint `gorm:"not null;index"`
	BranchID   uint `gorm:"not null;index"`

	EmployeeNo    string       `gorm:"type:varchar(50);not null;uniqueIndex"`
	Designation   string       `gorm:"type:varchar(150)"`
	Qualification string       `gorm:"type:varchar(500);not null"`
	JoiningDate   time.Time    `gorm:"not null"`
	Status        TeacherStatus `gorm:"type:varchar(20);not null;default:'ACTIVE';check:status IN ('ACTIVE','INACTIVE')"`
	
	// Relationships
	User *users.User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Coaching *coaching.Coaching `gorm:"forignKey:CoachingID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Branch *branches.Branch `gorm:"forignKey:BranchID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type TeacherSubject struct {
	gorm.Model

	TeacherID uint `gorm:"not null;index;uniqueIndex:idx_teacher_subject"`
	SubjectID uint `gorm:"not null;index;uniqueIndex:idx_teacher_subject"`

	// Relationships
	Teacher *Teacher `gorm:"foreignKey:TeacherID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Subject *subjects.Subject `gorm:"foreignKey:SubjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}