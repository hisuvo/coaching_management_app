package dto

import "time"

type BranchResponse struct {
	ID uint `json:"id"`
	CoachingID uint `json:"coaching_id"`
	Name string `json:"name"`
	Code string `json:"code"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Address string `json:"address"`
	City string `json:"city"`
	Country string `json:"country"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}