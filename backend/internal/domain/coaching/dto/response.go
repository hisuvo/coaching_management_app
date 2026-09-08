package dto

import (
	"coaching_backend/internal/domain/teachers/dto"
	"time"
)

type CoachingResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Domain    string `json:"domain"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CoachingWithTeachersResponse struct {
	ID        uint                  `json:"id"`
	Name      string                `json:"name"`
	Email     string                `json:"email"`
	Phone     string                `json:"phone"`
	Domain    string                `json:"domain"`
	Teachers  []dto.TeacherResponse `json:"teachers,omitempty"`
	CreatedAt time.Time                `json:"created_at"`
	UpdatedAt time.Time                `json:"updated_at"`
}