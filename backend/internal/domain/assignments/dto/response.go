package dto

import "time"

type AssignmentResponse struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Description *string    `json:"description,omitempty"`
	SubjectID   uint       `json:"subject_id"`
	CoachingID  uint       `json:"coaching_id"`
	CreatedBy   uint       `json:"created_by"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}