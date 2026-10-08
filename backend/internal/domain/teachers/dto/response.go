package dto

import (
	branchdto "coaching_backend/internal/domain/branches/dto"
	"coaching_backend/internal/domain/subjects/dto"
	usersdto "coaching_backend/internal/domain/users/dto"

	"time"
)

type TeacherResponse struct {
	ID          uint         `json:"id"`
	CoachingID  uint         `json:"coaching_id"`
	UserID      uint         `json:"user_id"`
	BranchID    uint         `json:"branch_id"`
	EmployeeNo  string       `json:"employee_no"`
	Designation string       `json:"designation"`
	JoiningDate *time.Time   `json:"joining_date,omitempty"`
	Status      string       `json:"status"`

	User   *usersdto.UserResponse   `json:"user,omitempty"`
	Branch *branchdto.BranchResponse `json:"branch,omitempty"`
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