package subjects

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrSubjectNotFound = errors.New("Not found subjects")
)

type Repoistory interface{
	Create(subject *Subject) error
	GetAll() ([]*Subject, error)
}

type repository struct{
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repoistory {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(subject *Subject)error{
	response := r.db.Create(subject).Error
	return response
}

func (r *repository) GetAll()([]*Subject, error) {
	var response []*Subject

	err := r.db.Find(&response).Error

	if err != nil {
		return nil, ErrSubjectNotFound
	}

	return response, nil
}