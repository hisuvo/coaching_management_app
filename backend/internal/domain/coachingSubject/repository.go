package coachingsubject

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, coachingSubject *CoachingSubject) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository (db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, coachingsubject *CoachingSubject) error {
	return r.db.Create(&coachingsubject).Error
}


