package dto

import "time"

type BatchResponse struct {
	Name    string `json:"name"`
	Code    string `json:"code"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Address string `json:"address"`
	City     string `json:"city"`
	Division string `json:"division"`
	Status string `json:"status"`
	TimeZone string `json:"time_zone"`
	OpeningTime *time.Time `json:"opening_time"`
	ClosingTime *time.Time `json:"closing_time"`
}