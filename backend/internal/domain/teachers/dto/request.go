package dto

import (
	// "coaching_backend/internal/domain/teachers"
	"time"
)

// CreateTeacherRequest represents the payload required to create a teacher.
type CreateTeacherRequest struct {
	UserID        uint         				`json:"user_id" validate:"required"`
	BranchID      uint         				`json:"branch_id" validate:"required"`
	EmployeeNo    string       				`json:"employee_no" validate:"required,max=50"`
	Designation   string       				`json:"designation,omitempty" validate:"max=150"`
	Qualification string       				`json:"qualification" validate:"required,max=500"`
	JoiningDate   time.Time    				`json:"joining_date"`
	Status        string	`json:"status,omitempty"`
}

// UpdateTeacherRequest represents the fields that can be updated.
type UpdateTeacherRequest struct {
	BranchID      *uint         			`json:"branch_id,omitempty"`
	Designation   *string       			`json:"designation,omitempty" validate:"omitempty,max=150"`
	Qualification *string       			`json:"qualification,omitempty" validate:"omitempty,max=500"`
	JoiningDate   *time.Time    			`json:"joining_date,omitempty"`
	Status        *string					`json:"status,omitempty"`
}

type CreateTeacherSubjectRequest struct {
	TeacherID uint `json:"teacher_id" validate:"required"`
	SubjectID uint `json:"subject_id" validate:"required"`
}

type UpdateTeacherSubjectRequest struct {
	TeacherID uint `json:"teacher_id" validate:"required"`
	SubjectID uint `json:"subject_id" validate:"required"`
}