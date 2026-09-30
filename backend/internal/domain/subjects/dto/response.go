package dto

import "time"

type SubjectResponse struct {
	Id          string		`json:"id"`
	Name        string		`json:"name"`
	Code        string		`json:"code"`
	Description *string		`json:"description,omitempty"`
	CreatedAt   time.Time	`json:"createdAt"`
	UpdatedAt 	time.Time	`json:"updatedAt"`
}
