package dto

type CreateAssignmentRequest struct {
	Title       string  `json:"title" validate:"required,min=3,max=200"`
	Description *string `json:"description,omitempty"`
	SubjectID   uint    `json:"subject_id" validate:"required"`
	CoachingID  uint    `json:"coaching_id" validate:"required"`
	DueDate     *string `json:"due_date,omitempty"`
	Status      string  `json:"status,omitempty"`
}

type UpdateAssignmentRequest struct {
	Title       *string `json:"title,omitempty" validate:"omitempty,min=3,max=200"`
	Description *string `json:"description,omitempty"`
	SubjectID   *uint   `json:"subject_id,omitempty"`
	DueDate     *string `json:"due_date,omitempty"`
	Status      *string `json:"status,omitempty"`
}