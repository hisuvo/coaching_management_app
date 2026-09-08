package dto

type CreateTeacherRequest struct {
	CoachingID  uint   `json:"coaching_id" validate:"required,gt=0"`
	Name        string `json:"name" validate:"required,min=2,max=100"`
	Email       string `json:"email" validate:"required,email"`
	Phone       string `json:"phone" validate:"required,min=11,max=20"`
	Password    string `json:"password" validate:"required,min=6"`
	Subject     string `json:"subject" validate:"required,max=100"`
	Designation string `json:"designation" validate:"omitempty,max=100"`
	Address     string `json:"address" validate:"omitempty,max=255"`
	IsActive    *bool  `json:"is_active"`
}

type UpdateTeacherRequest struct {
	Name        string `json:"name" validate:"omitempty,min=2,max=100"`
	Email       string `json:"email" validate:"omitempty,email"`
	Phone       string `json:"phone" validate:"omitempty,min=11,max=20"`
	Password    string `json:"password" validate:"omitempty,min=6"`
	Subject     string `json:"subject" validate:"omitempty,max=100"`
	Designation string `json:"designation" validate:"omitempty,max=100"`
	Address     string `json:"address" validate:"omitempty,max=255"`
	IsActive    *bool  `json:"is_active"`
}