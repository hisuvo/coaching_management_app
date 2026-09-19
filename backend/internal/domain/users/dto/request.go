package dto

import "time"

type CreateUserRequest struct {
	Name         string     `json:"name" validate:"required, min=2, max=100"`
	Email        string     `json:"email" validate:"required, email"`
	PasswordHash string     `json:"password_hash" validate:"required"`
	Role         string     `json:"role,omitempty"`
	Phone        string     `json:"phone"`
	LastLoginAt  *time.Time `json:"last_login_at"`
}

type UpdateUserRequest struct {
	Name   string `json:"name,omitempty" validate:"omitempty, min=2, max=100"`
	Email  string `json:"email,omitempty" validate:"omitempty, email"`
	Role   string `json:"role,omitempty"`
	Status string `json:"status,omitempty"`
}