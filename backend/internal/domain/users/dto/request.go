package dto

type CreateUserRequest struct {
	Name       string `json:"name" validate:"required,min=2,max=100"`
	Email      string `json:"email" validate:"required,email,max=255"`
	Phone      string `json:"phone" validate:"omitempty,max=20"`
	Password   string `json:"password" validate:"required,min=8,max=72"`
	CoachingID *uint  `json:"coaching_id" validate:"omitempty"`
}

type UpdateUserRequest struct {
	Name   *string `json:"name" validate:"omitempty,min=2,max=100"`
	Phone  *string `json:"phone" validate:"omitempty,max=20"`
	Status *string `json:"status" validate:"omitempty,oneof=ACTIVE INACTIVE SUSPENDED"`
}

type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required,min=8,max=72"`
}