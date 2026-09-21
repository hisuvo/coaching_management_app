package users

import (
	"time"

	"gorm.io/gorm"
)


type UserRole string

const (
	RoleSuperAdmin UserRole = "SUPER_ADMIN"
	RoleAdmin      UserRole = "ADMIN"
	RoleTeacher    UserRole = "TEACHER"
	RoleStudent    UserRole = "STUDENT"
	RoleMember    	UserRole = "MEMBER"
)

type UserStatus string

const (
	StatusActive  UserStatus = "ACTIVE"
	StatusInactive UserStatus = "INACTIVE"
	StatusBlocked  UserStatus = "BLOCKED"
)

type User struct {
	gorm.Model

	CoachingID *uint `gorm:"index"`
	Name string `gorm:"size:100;not null"`
	Email string `gorm:"size:255;uniqueIndex;not null"`
	PasswordHash  string `gorm:"not null" json:"-"`
	Role UserRole `gorm:"type:varchar(30);not null;index"`
	Status UserStatus `gorm:"varchar(30);not null; default:'ACTIVE';index"`
	Phone string `gorm:"type:varchar(20)"`
	LastLoginAt *time.Time
}
