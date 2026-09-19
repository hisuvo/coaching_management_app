package dto

import "time"

type UserResponse struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Status    string    `json:"status"`
	Phone        string `json:"phone"`
    LastLoginAt  *time.Time
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserListResponse struct {
	Users []UserResponse `json:"users"`
	Total int64 		 `json:"total"`
}