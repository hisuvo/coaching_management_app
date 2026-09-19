package dto

import (
	"coaching_backend/internal/domain/exams"
	"time"
)

type ExamResponse struct {
	ID uint `json:"id"`

	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	ExamType    exams.ExamType  `json:"exam_type"`
	Status      exams.ExamStatus `json:"status"`

	BranchID  uint `json:"branch_id"`
	BatchID   uint `json:"batch_id"`
	SubjectID uint `json:"subject_id"`
	CreatedBy uint `json:"created_by"`

	ExamDate       time.Time `json:"exam_date"`
	StartTime      time.Time `json:"start_time"`
	EndTime        time.Time `json:"end_time"`
	DurationMinute uint      `json:"duration_minute"`

	TotalMarks float64 `json:"total_marks"`
	PassMarks  float64 `json:"pass_marks"`

	RoomNumber   string `json:"room_number,omitempty"`
	Instructions string `json:"instructions,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}