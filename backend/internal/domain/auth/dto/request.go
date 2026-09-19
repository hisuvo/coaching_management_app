package dto

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken,omitempty"`
}

type LogoutRequest struct {
	RefreshToken string `json:"refreshToken,omitempty"`
}