package database

import (
	"coaching_backend/internal/domain/assignments"
	"coaching_backend/internal/domain/auth"
	"coaching_backend/internal/domain/coaching"
	"coaching_backend/internal/domain/students"
	"coaching_backend/internal/domain/subjects"
	"coaching_backend/internal/domain/submissions"
	"coaching_backend/internal/domain/users"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&auth.AuthSession{},
		&users.User{},
		&subjects.Subject{},
		&students.Student{},
		&coaching.Coaching{},
		&submissions.Submission{},
		&assignments.Assignment{},
	)
}