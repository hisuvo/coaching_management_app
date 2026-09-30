package subjects

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

var (
	ErrSubjectNotFound = errors.New("Not found subjects")
	ErrSubjectAlreadyExists = errors.New("Subject already exists")
)

type Repoistory interface{
	Create(ctx context.Context, subject *Subject) error
	GetAll() ([]*Subject, error)
	GetById(subjectId string)(*Subject, error)
	FindByCode(ctx context.Context, code string) (*Subject, error)
	Update(subjectId string, subject *Subject) (*Subject, error)
	Delete(subjectId string) (*Subject, error)
}

type repository struct{
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repoistory {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, subject *Subject)error{
	return r.db.WithContext(ctx).Create(subject).Error
}

func (r *repository) GetAll()([]*Subject, error) {
	var response []*Subject

	err := r.db.Find(&response).Error

	if err != nil {
		return nil, ErrSubjectNotFound
	}

	return response, nil
}

func (r *repository) GetById(subjectId string)(*Subject, error){
	var subject *Subject

	err := r.db.Where("id = ?", subjectId).First(&subject).Error

	if err != nil {
		return nil, err
	}

	return subject, nil
}

func (r *repository) FindByCode(ctx context.Context, code string) (*Subject, error) {
    var subject Subject

    if err := r.db.WithContext(ctx).Where("code = ?", code).First(&subject).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, ErrSubjectNotFound
        }
        return nil, err
    }

    return &subject, nil
}

func (r *repository) Update(subjectId string, subject *Subject) (*Subject, error) {
	var existing *Subject

	err := r.db.First(&existing, subjectId).Error
	
	if err != nil {
		return nil, err
	}


	err = r.db.Model(&existing).Updates(subject).Error

	if err != nil {
		return nil, err
	}

	return existing, nil
}

func (r *repository) Delete(subjectId string) (*Subject, error) {
	var subject Subject

	// find first subject
	if err := r.db.Where("id = ?", subjectId).First(&subject).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrSubjectNotFound
		}
		return nil, err
	} 

	// soft delete
	result := r.db.Delete(&subject)

	if result.Error != nil {
		return nil, result.Error
	}

	return &subject, nil
}
