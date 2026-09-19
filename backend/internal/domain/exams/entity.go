package exams

import (
	"time"

	"gorm.io/gorm"
)

type ExamType string

const (
	ExamTypeMonthly    ExamType = "monthly"
	ExamTypeWeekly     ExamType = "weekly"
	ExamTypeModelTest  ExamType = "model_test"
	ExamTypeMockTest   ExamType = "mock_test"
	ExamTypeFinal      ExamType = "final"
	ExamTypeHalfYearly ExamType = "half_yearly"
	ExamTypeOther      ExamType = "other"
)

type ExamStatus string

const (
	ExamStatusDraft     ExamStatus = "draft"
	ExamStatusScheduled ExamStatus = "scheduled"
	ExamStatusOngoing   ExamStatus = "ongoing"
	ExamStatusCompleted ExamStatus = "completed"
	ExamStatusCancelled ExamStatus = "cancelled"
)

type Exam struct {
	ID uint `json:"id" gorm:"primaryKey"`

	// Basic information
	Title       string    `json:"title" gorm:"type:varchar(150);not null;index"`
	Description string    `json:"description,omitempty" gorm:"type:text"`
	ExamType    ExamType  `json:"exam_type" gorm:"type:varchar(30);not null;index"`
	Status      ExamStatus `json:"status" gorm:"type:varchar(20);not null;default:'draft';index"`

	// Academic relationships
	BranchID  uint `json:"branch_id" gorm:"not null;index"`
	BatchID   uint `json:"batch_id" gorm:"not null;index"`
	SubjectID uint `json:"subject_id" gorm:"not null;index"`

	// Exam creator / examiner
	CreatedBy uint `json:"created_by" gorm:"not null;index"`

	// Exam configuration
	ExamDate       time.Time `json:"exam_date" gorm:"not null;index"`
	StartTime      time.Time `json:"start_time" gorm:"not null"`
	EndTime        time.Time `json:"end_time" gorm:"not null"`
	DurationMinute uint      `json:"duration_minute" gorm:"not null"`
	TotalMarks     float64   `json:"total_marks" gorm:"type:numeric(8,2);not null"`
	PassMarks      float64   `json:"pass_marks" gorm:"type:numeric(8,2);not null"`

	// Optional information
	RoomNumber string `json:"room_number,omitempty" gorm:"type:varchar(50)"`
	Instructions string `json:"instructions,omitempty" gorm:"type:text"`

	// Audit fields
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}