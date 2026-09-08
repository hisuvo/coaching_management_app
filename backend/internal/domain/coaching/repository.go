package coaching

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Create(coaching *Coaching) error
	GetById(id string) (*Coaching, error)
	GetAll() ([]*Coaching, error)
}

type repository struct {
	db *gorm.DB
}

var (
	ErrNotFoundCoaching = errors.New("Coaching Not Found")
)

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(coaching *Coaching) error {
	return r.db.Create(coaching).Error
}

func (r *repository) GetById(id string) (*Coaching, error){
	var coaching Coaching

	if err := r.db.Where("id = ?", id).First(&coaching).Error; err != nil {
		return nil, ErrNotFoundCoaching
	}

	return &coaching, nil
}

func (r *repository) GetAll() ([]*Coaching, error){
	var coachings []*Coaching

	if err := r.db.Find(&coachings).Error; err != nil {
		return nil, ErrNotFoundCoaching
	}

	return coachings, nil
}