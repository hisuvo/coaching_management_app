package students

import "gorm.io/gorm"

type Repository interface {
	Create(student *Student) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(student *Student) error {
	return r.db.Create(student).Error
}