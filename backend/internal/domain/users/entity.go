package users

import "gorm.io/gorm"


type Role string
type UserStatus string

type User struct {
	gorm.Model
	Name string `gorm:"size:100;not null" json:"name"`
	Email string `gorm:"size:255;uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"-"`
	Role Role `gorm:"size:30;not null; default:user" json:"role"`
	Status UserStatus `gorm:"size:50;not null; default:active" json:"status"`
}


const (
	RoleSuperAdmin Role = "super_admin"
	RoleAdmin Role = "admin"
	RoleStudent Role = "student"
	RoleTeacher Role = "teacher"
	RoleGuardian Role = "guardian"
	RoleUser  Role = "user"
)


const (
	UserStatusActive UserStatus = "active"
	UserStatusBlock UserStatus = "block"
)