package dto

import "time"


type CreateBatchRequest struct {
	Name        string     `json:"name" validate:"required,min=2,max=150"`
	Code        string     `json:"code" validate:"required,min=2,max=50"`
	Phone       string     `json:"phone,omitempty" validate:"omitempty,max=20"`
	Email       string     `json:"email,omitempty" validate:"omitempty,email,max=150"`
	Address     string     `json:"address,omitempty"`
	City        string     `json:"city,omitempty" validate:"omitempty,max=100"`
	Division    string     `json:"division,omitempty" validate:"omitempty,max=100"`
	Status      string     `json:"status,omitempty" validate:"omitempty,oneof=ACTIVE INACTIVE"`
	TimeZone    string     `json:"time_zone,omitempty" validate:"omitempty,max=100"`
	OpeningTime *time.Time `json:"opening_time,omitempty"`
	ClosingTime *time.Time `json:"closing_time,omitempty"`
}

type UpdateBatchRequest struct {
	Name        *string    `json:"name,omitempty" validate:"omitempty,min=2,max=150"`
	Code        *string    `json:"code,omitempty" validate:"omitempty,min=2,max=50"`
	Phone       *string    `json:"phone,omitempty" validate:"omitempty,max=20"`
	Email       *string    `json:"email,omitempty" validate:"omitempty,email,max=150"`
	Address     *string    `json:"address,omitempty"`
	City        *string    `json:"city,omitempty" validate:"omitempty,max=100"`
	Division    *string    `json:"division,omitempty" validate:"omitempty,max=100"`
	Status      *string    `json:"status,omitempty" validate:"omitempty,oneof=ACTIVE INACTIVE"`
	TimeZone    *string    `json:"time_zone,omitempty" validate:"omitempty,max=100"`
	OpeningTime *time.Time `json:"opening_time,omitempty"`
	ClosingTime *time.Time `json:"closing_time,omitempty"`
}