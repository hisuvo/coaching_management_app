package dto

import "time"

// SubmissionResponse represents the submission data
// returned to the client.
type SubmissionResponse struct {
	ID           uint       `json:"id"`
	AssignmentID uint       `json:"assignmentId"`
	StudentID    uint       `json:"studentId"`

	TextAnswer *string `json:"textAnswer,omitempty"`
	FileURL    *string `json:"fileUrl,omitempty"`

	SubmittedAt *time.Time `json:"submittedAt,omitempty"`

	Status string `json:"status"`

	ObtainedMarks *float64 `json:"obtainedMarks,omitempty"`
	Feedback      *string  `json:"feedback,omitempty"`

	ReviewedBy *uint      `json:"reviewedBy,omitempty"`
	ReviewedAt *time.Time `json:"reviewedAt,omitempty"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type SubmissionListResponse struct {

	// Data contains submission records.
	Data []*SubmissionResponse `json:"data"`
	
	// Page represents current page.
	Page int `json:"page"`

	// Limit represents current page size.
	Limit int `json:"limit"`

	// Total represents total matching records.
	Total int64 `json:"total"`

	// TotalPages represents total available pages.
	TotalPages int `json:"totalPages"`
}