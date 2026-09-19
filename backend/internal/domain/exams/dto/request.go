package dto

import "coaching_backend/internal/domain/exams"

type CreateExamRequest struct {
	Title       string `json:"title" validate:"required,min=3,max=150"`
	Description string `json:"description,omitempty" validate:"omitempty,max=2000"`

	ExamType exams.ExamType `json:"exam_type" validate:"required,oneof=monthly weekly model_test mock_test final half_yearly other"`

	BranchID  uint `json:"branch_id" validate:"required,gt=0"`
	BatchID   uint `json:"batch_id" validate:"required,gt=0"`
	SubjectID uint `json:"subject_id" validate:"required,gt=0"`

	ExamDate       string `json:"exam_date" validate:"required"`
	StartTime      string `json:"start_time" validate:"required"`
	EndTime        string `json:"end_time" validate:"required"`
	DurationMinute uint   `json:"duration_minute" validate:"required,gt=0,lte=600"`

	TotalMarks float64 `json:"total_marks" validate:"required,gt=0"`
	PassMarks  float64 `json:"pass_marks" validate:"required,gte=0"`

	RoomNumber   string `json:"room_number,omitempty" validate:"omitempty,max=50"`
	Instructions string `json:"instructions,omitempty" validate:"omitempty,max=5000"`
}

type UpdateExamRequest struct {
	Title        *string    `json:"title,omitempty" validate:"omitempty,min=3,max=150"`
	Description  *string    `json:"description,omitempty" validate:"omitempty,max=2000"`
	ExamType     *exams.ExamType  `json:"exam_type,omitempty" validate:"omitempty,oneof=monthly weekly model_test mock_test final half_yearly other"`
	Status       *exams.ExamStatus `json:"status,omitempty" validate:"omitempty,oneof=draft scheduled ongoing completed cancelled"`

	BranchID  *uint `json:"branch_id,omitempty" validate:"omitempty,gt=0"`
	BatchID   *uint `json:"batch_id,omitempty" validate:"omitempty,gt=0"`
	SubjectID *uint `json:"subject_id,omitempty" validate:"omitempty,gt=0"`

	ExamDate       *string  `json:"exam_date,omitempty"`
	StartTime      *string  `json:"start_time,omitempty"`
	EndTime        *string  `json:"end_time,omitempty"`
	DurationMinute *uint    `json:"duration_minute,omitempty" validate:"omitempty,gt=0,lte=600"`

	TotalMarks *float64 `json:"total_marks,omitempty" validate:"omitempty,gt=0"`
	PassMarks  *float64 `json:"pass_marks,omitempty" validate:"omitempty,gte=0"`

	RoomNumber   *string `json:"room_number,omitempty" validate:"omitempty,max=50"`
	Instructions *string `json:"instructions,omitempty" validate:"omitempty,max=5000"`
}