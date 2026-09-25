package database

import (
	"coaching_backend/internal/domain/assignments"
	"coaching_backend/internal/domain/auth"
	"coaching_backend/internal/domain/branches"
	"coaching_backend/internal/domain/coaching"
	"coaching_backend/internal/domain/students"
	"coaching_backend/internal/domain/subjects"
	"coaching_backend/internal/domain/submissions"
	"coaching_backend/internal/domain/users"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&users.User{},
		&auth.AuthSession{},
		&coaching.Coaching{},
		&branches.Branch{},
		&subjects.Subject{},
		&students.Student{},
		&submissions.Submission{},
		&assignments.Assignment{},
	)
}