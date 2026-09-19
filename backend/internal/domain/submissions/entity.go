package submissions

import (
	"time"

	"gorm.io/gorm"
)

type SubmissionStatus string

const (
    SubmissionStatusSubmitted SubmissionStatus = "SUBMITTED"
    SubmissionStatusReviewed SubmissionStatus = "REVIEWED"
    SubmissionStatusReturned SubmissionStatus = "RETURNED"
    SubmissionStatusLate SubmissionStatus = "LATE"
)

type Submission struct {
	gorm.Model

    CoaschingId  uint `json:"coaschingId" gorm:"not null;index"`
	AssignmentID uint `json:"assignmentId" gorm:"not null;index"`
    StudentID    uint `json:"studentId" gorm:"not null;index"`

    FileURL       string  `json:"fileUrl" gorm:"type:text"`
    TextAnswer    *string `json:"textAnswer" gorm:"type:text"`
    SubmittedAt   *time.Time `json:"submittedAt"`
    Status        string `json:"status" gorm:"type:varchar(30);not null;index"`
    ObtainedMarks *float64 `json:"obtainedMarks"`
    Feedback      *string `json:"feedback" gorm:"type:text"`
    ReviewedBy      *uint `json:"reviewedBy"`
    ReviewedAt      *time.Time `json:"reviewedAt"`
}