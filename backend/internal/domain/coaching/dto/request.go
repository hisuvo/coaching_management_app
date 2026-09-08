package dto

type CreateCoachingRequest struct {
	Name   string `json:"name" validate:"required,min=2,max=150"`
	Email  string `json:"email" validate:"required,email"`
	Phone  string `json:"phone" validate:"required,min=11,max=20"`
	Domain string `json:"domain" validate:"required,max=255"`
}

type UpdateCoachingRequest struct {
	Name   string `json:"name" validate:"omitempty,min=2,max=150"`
	Email  string `json:"email" validate:"omitempty,email"`
	Phone  string `json:"phone" validate:"omitempty,min=11,max=20"`
	Domain string `json:"domain" validate:"omitempty,max=255"`
}