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
	FindById(subjectId string)(*Subject, error)
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

func (r *repository) FindById(subjectId string)(*Subject, error){
	var subject *Subject

	err := r.db.Where("id = ?", subjectId).First(&subject).Error

	if err != nil {
		return nil, err
	}

	return subject, nil
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

