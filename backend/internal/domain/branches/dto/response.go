package dto

import (
	"time"
)

type BranchResponse struct {
	ID uint `json:"id,omitempty"`
	CoachingID uint `json:"coaching_id,omitempty"`
	Name string `json:"name"`
	Code string `json:"code"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Address string `json:"address,omitempty"`
	Status string `json:"status"`

	Admin *AdminResponse `json:"amdin,omitempty"`
}

type BranchDetailsResponse struct {
	ID uint `json:"id"`
	CoachingID uint `json:"coaching_id"`
	Name string `json:"name"`
	Code string `json:"code"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Address string `json:"address"`
	Status string `json:"status"`

	Admin *AdminResponse `json:"amdin,omitempty"`
	Students int64
	Teachers int64

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AdminResponse struct {
	ID    uint   `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Role  string `json:"role,omitempty"`
}