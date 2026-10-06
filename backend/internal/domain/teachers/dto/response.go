package dto

import (
	"coaching_backend/internal/domain/subjects/dto"
	"time"
)

type TeacherResponse struct {
	UserID        uint       `json:"user_id"`
	BranchID      uint       `json:"branch_id"`
	EmployeeNo    string     `json:"employee_no"`
	Designation   string     `json:"designation,omitempty"`
	Qualification string     `json:"qualification"`
	JoiningDate   time.Time  `json:"joining_date"`
	Status        string	 `json:"status"`
	CreatedAt	 time.Time	 `json:"created_at"`
	UpdatedAt	 time.Time	 `json:"updated_at"`
}

type TeacherSubjectResponse struct {
	ID        uint `json:"id"`
	TeacherID uint `json:"teacher_id"`
	SubjectID uint `json:"subject_id"`

	Teacher *TeacherResponse     `json:"teacher,omitempty"`
	Subject *dto.SubjectResponse `json:"subject,omitempty"`
	CreatedAt	 time.Time		 `json:"created_at"`
	UpdatedAt	 time.Time		 `json:"updated_at"`
}