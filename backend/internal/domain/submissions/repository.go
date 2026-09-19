package submissions

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Create(submission *Submission) ( error)
	FindByStudentAndAssignment(studentId string, assignmentId string, coachingId string) (*Submission, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository (db *gorm.DB) Repository {
	return &repository{
		db,
	}
}

func (r *repository) Create(submission *Submission) error{
	// Insert submission into database
	return r.db.Create(submission).Error
}

func (r *repository) FindByStudentAndAssignment(studentId string, assignmentId string, coachingId string) (*Submission, error) {
	var submission Submission

	err := r.db.Where("studentId = ? AND assignmentId = ? AND coachingId = ?", studentId, assignmentId, coachingId).First(submission).Error

	if err != nil {
		return nil, err
	}

	return &submission, nil
}

func IsNotFound(err error) bool {
	// Compare error against GORM's not-found error.
	return errors.Is(err, gorm.ErrRecordNotFound)
}