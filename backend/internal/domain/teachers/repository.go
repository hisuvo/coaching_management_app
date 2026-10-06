package teachers

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

var (
	ErrTeacherNotFound  = errors.New("teacher not found")
	ErrEmployeeNoExists = errors.New("employee number already exists")
)

type Repository interface{
	Create(ctx context.Context, teacher *Teacher) (*Teacher, error)
	GetByID(ctx context.Context, coachingID uint, id uint) (*Teacher, error) 
	GetAll(ctx context.Context, coachingID uint) ([]*Teacher, error)
	Update(ctx context.Context, coachingID uint, id uint, teacher *Teacher) (*Teacher, error)
	Delete(ctx context.Context, coachingID uint, id uint) error
	FindByEmployeeNo(ctx context.Context, coachingID uint, employeeNo string) (*Teacher, error)
}

// WHERE id = ? AND coaching_id = ?

type repository struct {
	db *gorm.DB
}

func NewRegister(db *gorm.DB) Repository{
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context,teacher *Teacher) (*Teacher, error) {

	if err := r.db.WithContext(ctx).Create(teacher).Error; err != nil {return nil, err}

	return teacher, nil
}

func (r *repository) GetByID(ctx context.Context,coachingID uint,id uint,) (*Teacher, error) {

	var teacher Teacher

	err := r.db.WithContext(ctx).Where("coaching_id = ? AND id = ?", coachingID, id).First(&teacher).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, ErrTeacherNotFound
		}
		return nil, err
	}

	return &teacher, nil
}

func (r *repository) GetAll(ctx context.Context,coachingID uint) ([]*Teacher, error) {

	var teachers []*Teacher

	err := r.db.WithContext(ctx).Where("coaching_id = ?", coachingID).Find(&teachers).Error

	if err != nil {
		return nil, err
	}

	return teachers, nil
}

func (r *repository) Update(ctx context.Context,coachingID uint,id uint,teacher *Teacher) (*Teacher, error) {

	var existingTeacher Teacher

	err := r.db.WithContext(ctx).Where("coaching_id = ? AND id = ?", coachingID, id).First(&existingTeacher).Error

	if err != nil {
		return nil, err
	}

	if err := r.db.WithContext(ctx).Model(&existingTeacher).Updates(teacher).Error; err != nil {
		return nil, err
	}

	return &existingTeacher, nil
}

func (r *repository) Delete(ctx context.Context,coachingID uint,id uint) error {

	result := r.db.WithContext(ctx).Model(&Teacher{}).Where("coaching_id = ? AND id = ?", coachingID, id).Delete(&Teacher{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *repository) FindByEmployeeNo(ctx context.Context,coachingID uint,employeeNo string) (*Teacher, error) {

	var teacher Teacher

	err := r.db.WithContext(ctx).Where("coachingID = ? AND employee_no = ?",coachingID,employeeNo).First(&teacher).Error

	if err != nil {
		return nil, err
	}

	return &teacher, nil
}