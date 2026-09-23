package dto

type CreateCoachingRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=150"`
	Email    string `json:"email" validate:"required,email,max=150"`
	Phone    string `json:"phone" validate:"required,max=20"`
	LogoURL  string `json:"logo_url,omitempty"`
	Address  string `json:"address,omitempty"`
	TimeZone string `json:"time_zone,omitempty"`
}

type UpdateCoachingRequest struct {
	Name     *string `json:"name,omitempty" validate:"omitempty,min=3,max=150"`
	Email    *string `json:"email,omitempty" validate:"omitempty,email,max=150"`
	Phone    *string `json:"phone,omitempty" validate:"omitempty,max=20"`
	LogoURL  *string `json:"logo_url,omitempty"`
	Address  *string `json:"address,omitempty"`
	TimeZone *string `json:"time_zone,omitempty"`
	Status   *string `json:"status,omitempty"`
}