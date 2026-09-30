package dto

import "time"

type CoachingResponse struct {
	ID       uint           `json:"id"`
	Name     string         `json:"name"`
	Email    string         `json:"email"`
	Phone    string         `json:"phone"`
	Domain   string         `json:"domain"`
	Slug     string         `json:"slug"`
	LogoURL  string         `json:"logo_url,omitempty"`
	Address  string         `json:"address,omitempty"`
	TimeZone string         `json:"time_zone"`
	Status   string         `json:"status"`
	Admin    *AdminResponse `json:"admin,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CoachingDetailResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Domain   string `json:"domain"`
	Slug     string `json:"slug"`
	LogoURL  string `json:"logo_url,omitempty"`
	Address  string `json:"address,omitempty"`
	TimeZone string `json:"time_zone"`
	Status   string `json:"status"`

	Admin    *AdminResponse `json:"admin,omitempty"`
	Branches int64          `json:"branches"`
	Teachers int64          `json:"teachers"`
	Students int64          `json:"students"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AdminResponse struct {
	ID    uint   `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Role  string `json:"role,omitempty"`
}