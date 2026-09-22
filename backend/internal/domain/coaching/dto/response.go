package dto

import (
	"coaching_backend/internal/domain/teachers/dto"
	"time"
)

type CoachingResponse struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Phone     string         `json:"phone"`
	Domain    string         `json:"domain"`
	Slug      string         `json:"slug"`
	LogoURL   string         `json:"logo_url,omitempty"`
	Address   string         `json:"address,omitempty"`
	TimeZone  string         `json:"time_zone"`
	Status    string         `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type CoachingWithTeachersResponse struct {
	CoachingResponse
	Teachers  []dto.TeacherResponse `json:"teachers,omitempty"`
}