package dto

// CreateSubmissionRequest represents the data required
// when a student submits an assignment.
type CreateSubmissionRequest struct {
	AssignmentID uint    `json:"assignmentId" validate:"required"`
	StudentID    uint 	 `json:"studentId" validate:"required"`
	TextAnswer   *string `json:"textAnswer,omitempty"`
	FileURL      *string `json:"fileUrl,omitempty"`
}

// UpdateSubmissionRequest represents the data that can be
// updated after a submission has been created.
type UpdateSubmissionRequest struct {
	TextAnswer *string `json:"textAnswer,omitempty"`
	FileURL    *string `json:"fileUrl,omitempty"`
	Status *string `json:"status,omitempty"`
	Marks *float64 `json:"marks,omitempty" validate:"omitempty,min=0"`
	Feedback *string `json:"feedback,omitempty"`
}

// ReviewSubmissionRequest represents teacher review data.
type ReviewSubmissionRequest struct {
	// Marks represents marks awarded to the student.
	Marks float64 `json:"marks" validate:"min=0"`
	// Feedback contains teacher feedback.
	Feedback *string `json:"feedback,omitempty"`
}

type SubmissionQuery struct {
	// AssignmentID filters submissions by assignment.
	AssignmentID uint `query:"assignmentId"`
	// StudentID filters submissions by student.
	StudentID uint `query:"studentId"`
	// Status filters submissions by status.
	Status string `query:"status"`
	// Page represents the requested page number.
	Page int `query:"page"`
	// Limit represents the number of records per page.
	Limit int `query:"limit"`
}
