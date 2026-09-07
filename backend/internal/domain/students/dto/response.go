package dto

import "time"

type StudentResponse struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Class      string    `json:"class"`
	Session   string    `json:"secssion"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	Address    string    `json:"address"`
	SchoolName string    `json:"school_name"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
}