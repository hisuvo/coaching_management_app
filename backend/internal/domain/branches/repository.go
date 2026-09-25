package branches

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

var (
	ErrBranchNotFound      = errors.New("branch not found by this code")
	ErrBranchAlreadyExists = errors.New("branch already exists")
	ErrBranchCreationFailed  = errors.New("branch creation failed")
	ErrBranchUpdateFailed  = errors.New("branch update failed")
	ErrBranchDeleteFailed  = errors.New("branch delete failed")
	ErrBranchFindFailed    = errors.New("branch find failed")
	ErrBranchFindAllFailed = errors.New("branch find all failed")
)

type Repository interface {
	Create(ctx context.Context, branch *Branch) (*Branch, error)
	FindByID(ctx context.Context, id uint) (*Branch, error)
	FindByCode(ctx context.Context, coachingID uint, code string) (*Branch, error)
	FindAllByCoachingID(ctx context.Context, coachingID uint) ([]*Branch, error)
	Update(ctx context.Context, id uint, branch *Branch) (*Branch, error)
	Delete(ctx context.Context, id uint) (*Branch, error)
	FindAll(ctx context.Context) ([]*Branch, error)
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

	if err := r.db.WithContext(ctx).Create(&branch).Error; err != nil {
		return nil, err
	}

	return branch, nil
}

func (r *repository) FindByID(ctx context.Context, id uint) (*Branch, error) {
	var branch Branch

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&branch).Error; err != nil {
		return nil, err
	}
	return &branch, nil
}

func (r *repository) Update(ctx context.Context, id uint, branch *Branch) (*Branch, error) {

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&branch).Error; err != nil {
		return nil, err
	}

	if branch.Name != "" {
		branch.Name = branch.Name
	}
	if branch.Code != "" {
		branch.Code = branch.Code
	}
	if branch.Phone != "" {
		branch.Phone = branch.Phone
	}
	if branch.Email != "" {
		branch.Email = branch.Email
	}
	if branch.Address != "" {
		branch.Address = branch.Address
	}
	if branch.City != "" {
		branch.City = branch.City
	}
	if branch.Country != "" {
		branch.Country = branch.Country
	}
	if branch.Status != "" {
		branch.Status = branch.Status
	}

	if err := r.db.WithContext(ctx).Save(&branch).Error; err != nil {
		return nil, err
	}
	return branch, nil
}

func (r *repository) Delete(ctx context.Context, id uint) (*Branch, error) {
	var branch Branch

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&branch).Error; err != nil {
		return nil, err
	}
	return &branch, nil
}

func (r *repository) FindAll(ctx context.Context) ([]*Branch, error) {
	var branches []*Branch
	if err := r.db.WithContext(ctx).Find(&branches).Error; err != nil {
		return nil, err
	}
	return branches, nil
}

func (r *repository) FindByCode(ctx context.Context, coachingID uint, code string) (*Branch, error) {
	var branch Branch
	if err := r.db.WithContext(ctx).Where("coaching_id = ? AND code = ?", coachingID, code).First(&branch).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, ErrBranchNotFound
		}
		return nil, err
	}
	return &branch, nil
}

func (r *repository) FindAllByCoachingID(ctx context.Context, coachingID uint) ([]*Branch, error) {
	var branches []*Branch
	if err := r.db.WithContext(ctx).Where("coaching_id = ?", coachingID).Find(&branches).Error; err != nil {
		return nil, err
	}
	return branches, nil
}