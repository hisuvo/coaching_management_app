package dto

type CreateCoachingRequest struct {
	Name     string `json:"name" validate:"required,max=150"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required"`
	LogoURL  string `json:"logo_url,omitempty"`
	Address  string `json:"address,omitempty"`
	TimeZone string `json:"time_zone,omitempty"`

	AdminName     string `json:"admin_name" validate:"required"`
	AdminEmail    string `json:"admin_email" validate:"required,email"`
	AdminPassword string `json:"admin_password" validate:"required,min=8"`
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