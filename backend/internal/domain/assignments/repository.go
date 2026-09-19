package assignments

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]*Assignment,error)
	Create(assignments *Assignment) error
	GetById(assignmentId string) (*Assignment, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository (db *gorm.DB) Repository{
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]*Assignment, error) {
	var assignment []*Assignment

	if err := r.db.Find(&assignment).Error; err != nil {
		return nil, err
	}

	return assignment, nil
}

func (r *repository) Create(assignment *Assignment) error {
	return r.db.Create(assignment).Error
}

func (r *repository) GetById(assignmentId string) (*Assignment, error) {
	var assignment Assignment

	if err := r.db.Where("id = ?", assignmentId).First(&assignment).Error; err != nil {
		return nil, err
	}

	return &assignment, nil
}