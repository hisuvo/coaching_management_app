package dto

import "time"

type UserResponse struct {
	ID         uint    `json:"id"`
	CoachingID *uint   `json:"coachingId,omitempty"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Role       string  `json:"role"`
	Status     string  `json:"status"`
	Phone      string  `json:"phone,omitempty"`
}

type LoginResponse struct {
	AccessToken string       `json:"accessToken"`
	TokenType   string       `json:"tokenType"`
	ExpiresIn   int64        `json:"expiresIn"`
	User        UserResponse `json:"user"`
}

type RefreshResponse struct {
	AccessToken string `json:"accessToken"`
	TokenType   string `json:"tokenType"`
	ExpiresIn   int64  `json:"expiresIn"`
}

type SessionResponse struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"userId"`
	IPAddress string    `json:"ipAddress,omitempty"`
	UserAgent string    `json:"userAgent,omitempty"`
	ExpiresAt time.Time `json:"expiresAt"`
	CreatedAt time.Time `json:"createdAt"`
}

// Todo: this struct is test porpose
type LoginResult struct {
	AccessToken string
	ExpiresIn int64
	RefreshToken string
	UserID uint
	CoachingID *uint
	Role string
	Name string
	Email string
}

type RefreshResult struct {
	AccessToken string
	ExpiresIn int64
	RefreshToken string
}