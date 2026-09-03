package dto

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required, min=2, max=100"`
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role,omitempty"`
}

type UpdateUserRequest struct {
	Name   string `json:"name,omitempty" validate:"omitempty, min=2, max=100"`
	Email  string `json:"email,omitempty" validate:"omitempty, email"`
	Role   string `json:"role,omitempty"`
	Status string `json:"status,omitempty"`
}