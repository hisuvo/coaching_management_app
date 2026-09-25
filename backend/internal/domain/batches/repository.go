package batches

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, brach *Branch)(*Branch, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, branch *Branch) (*Branch, error) {
	if err := r.db.WithContext(ctx).Create(branch).Error; err != nil {
		return nil, err
	}
	return  branch, nil
}