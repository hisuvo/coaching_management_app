package subjects

import "gorm.io/gorm"

type Repoistory interface{
	Create(subject *Subject) error
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
	user := r.db.Create(subject).Error
	return user
}