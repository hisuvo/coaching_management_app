package coachingsubject

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

var (
	ErrCoachingSubjectExist = errors.New("coaching subject already exists")
	ErrCoachingSubjectNotFound = errors.New("coaching subject not found")
)


type Repository interface {
	Create(ctx context.Context, coachingSubject *CoachingSubject) error
	FindCoachingSubjectByCoadeAndId(ctx context.Context, coaching_id uint, subject uint) (*CoachingSubject, error)
	GetOwnCoachingSubject(ctx context.Context, coaching_id *uint) ([]*CoachingSubject, error)
	GetSingleById(ctx context.Context, id uint) (*CoachingSubject, error)
	GetAll(ctx context.Context)([]*CoachingSubject, error)
	Upate(ctx context.Context, id uint, status *string) (*CoachingSubject, error)
	DeleteCoaching(ctx context.Context, id uint) error
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

func (r *repository) FindCoachingSubjectByCoadeAndId(ctx context.Context, coaching_id uint, subject_id uint) (*CoachingSubject, error){
	var coaching_subject CoachingSubject
	
	err := r.db.WithContext(ctx).Model(&CoachingSubject{}).Where("coaching_id = ? AND subject_id = ?", coaching_id, subject_id).First(&coaching_subject).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCoachingSubjectNotFound
		}

		return nil, err
	}

	return &coaching_subject, nil
}

func (r *repository) GetOwnCoachingSubject(ctx context.Context, coaching_id *uint) ([]*CoachingSubject, error){
	var OwnCoachingSubjects []*CoachingSubject

	if err := r.db.WithContext(ctx).Model(&CoachingSubject{}).Where("coaching_id = ?", coaching_id).Preload("Subject").Find(&OwnCoachingSubjects).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, ErrCoachingSubjectNotFound
		}

		return nil, err
	}

	if len(OwnCoachingSubjects) == 0 {
		return nil, ErrCoachingSubjectNotFound
	}

	return OwnCoachingSubjects, nil
}

func (r *repository) GetSingleById(ctx context.Context, id uint) (*CoachingSubject, error){
	var coaching_subject *CoachingSubject

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&coaching_subject).Error; err != nil {
		return nil, ErrCoachingSubjectNotFound
	}
	return coaching_subject, nil
}

func (r *repository) GetAll(ctx context.Context)([]*CoachingSubject, error) {
	var coaching_subjects []*CoachingSubject

	if err := r.db.WithContext(ctx).Preload("Subject").Find(&coaching_subjects).Error; err != nil {
		return nil, ErrCoachingSubjectNotFound
	}

	return coaching_subjects, nil
}

func (r *repository) Upate(ctx context.Context, id uint, status *string) (*CoachingSubject, error) {
	result := r.db.WithContext(ctx).Model(&CoachingSubject{}).Where("id = ?", id).Update("status", status);

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, ErrCoachingSubjectNotFound
	}

	// load and refetch update coaching subject data
	var coaching_subject CoachingSubject

	if err := r.db.WithContext(ctx).First(&coaching_subject, id).Error; err != nil {
		return nil, err
	}

	return &coaching_subject, nil
}

func (r *repository) DeleteCoaching(ctx context.Context, id uint) error {

	if err := r.db.WithContext(ctx).Model(&CoachingSubject{}).Where("id = ?", id).Update("deleted_at", time.Now()).Error; err != nil {
		return ErrCoachingSubjectNotFound
	} 

	return nil
}