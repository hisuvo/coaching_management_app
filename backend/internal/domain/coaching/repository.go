package coaching

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

var (
	ErrCoachingNotFound = errors.New("Coaching Not Found")
	ErrCoachingEmailExist = errors.New("Already user this email to create Caoching")
)


type Repository interface {
	Create(ctx context.Context, coaching *Coaching) error
	FindByEmail(ctx context.Context, email string) (*Coaching, error)
	GetById(ctx context.Context, id uint) (*Coaching, error)
	GetAll(ctx context.Context, ) ([]*Coaching, error)
	Update(ctx context.Context, id uint, coaching *Coaching) (*Coaching, error)
	Delete(ctx context.Context, id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context,coaching *Coaching) error {
	return r.db.Create(coaching).Error
}

func (r *repository) FindByEmail(ctx context.Context,email string) (*Coaching, error) {
	var coaching Coaching

	if err := r.db.Where("email = ?", email).First(&coaching).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, ErrCoachingNotFound
		}

		return nil, err
	}

	return &coaching,nil
}

func (r *repository) GetById(ctx context.Context,id uint) (*Coaching, error){
	var coaching Coaching

	if err := r.db.Where("id = ?", id).First(&coaching).Error; err != nil {
		if errors.Is(err,gorm.ErrRecordNotFound){
			return nil, ErrCoachingNotFound
		}
		return nil, err
	}

	return &coaching, nil
}

func (r *repository) GetAll(ctx context.Context,) ([]*Coaching, error){
	var coachings []*Coaching

	if err := r.db.Find(&coachings).Error; err != nil {
		return nil, ErrCoachingNotFound
	}

	return coachings, nil
}

func (r *repository) Update(ctx context.Context, id uint, coaching *Coaching) (*Coaching, error){
	var existing Coaching

	// First check whether the coaching exists.
	if err := r.db.WithContext(ctx).First(&existing, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, ErrCoachingNotFound
		}
		return nil, err
	}

	// Update only fields allowed by the service.
	updates := map[string]interface{}{
		"name":      coaching.Name,
		"email":     coaching.Email,
		"phone":     coaching.Phone,
		"logo_url":  coaching.LogoURL,
		"address":   coaching.Address,
		"time_zone": coaching.TimeZone,
		"status":    coaching.Status,
	}

	if err := r.db.WithContext(ctx).Model(&existing).Updates(updates).Error; err != nil {
		return nil, err
	}

	// Return the latest database record.
	if err := r.db.WithContext(ctx).First(&existing, id).Error; err != nil {
		return nil, err
	}

	return &existing, nil
}

func (r *repository) Delete(ctx context.Context,id uint) error{
	var coaching Coaching

	// Find this caochin is exists
	if err := r.db.WithContext(ctx).First(&coaching, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return ErrCoachingNotFound
		}

		return err
	}

	return r.db.WithContext(ctx).Delete(&coaching).Error
}