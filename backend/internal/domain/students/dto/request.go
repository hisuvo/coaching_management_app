package dto

type CreateStudentRequest struct {
	CoachingID string `json:"coaching_id" validate:"required,gt=0"`
	Name       string `json:"name" validate:"required,min=2,max=100"`
	Class      string `json:"class" validate:"required,max=50"`
	Session    string `json:"session" validate:"required,max=20"`
	Phone      string `json:"phone" validate:"required,min=11,max=20"`
	Email      string `json:"email" validate:"required,email,max=150"`
	Address    string `json:"address" validate:"required,max=500"`
	SchoolName string `json:"school_name" validate:"required,max=250"`
}

type UpdateStudentRequest struct {
	Name       *string `json:"name" validate:"omitempty,min=2,max=100"`
	Class      *string `json:"class" validate:"omitempty,max=50"`
	Session    *string `json:"session" validate:"omitempty,max=20"`
	Phone      *string `json:"phone" validate:"omitempty,max=20"`
	Email      *string `json:"email" validate:"omitempty,email,max=150"`
	Address    *string `json:"address" validate:"omitempty"`
	SchoolName *string `json:"school_name" validate:"omitempty,max=250"`
}
