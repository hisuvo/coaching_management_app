package users

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

var(
	 ErrDuplicateEmail = errors.New("Email already exists")
	 ErrEmailNotFound = errors.New("Email not found in record")
	 ErrUserNotRound = errors.New("User not found in record")
)

type Repository interface {
	Create (ctx context.Context,user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
	GetByID(ctx context.Context,id uint,) (*User, error)
	UpdateLastLogin(ctx context.Context,id uint,) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, user *User) error {

	err := r.db.WithContext(ctx).Create(user).Error

	if err != nil {

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return ErrDuplicateEmail
		}

		return err
	}
	return nil
}

func (r *repository) FindByEmail(ctx context.Context, email string) (*User, error){
	var user User

	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	return &user,nil
}

func(r *repository) GetByID(ctx context.Context, id uint) (*User, error){
	var user User
	err := r.db.WithContext(ctx).First(&user,id).Error

	if errors.Is(err, gorm.ErrRecordNotFound){
		return nil, ErrEmailNotFound
	}
	
	if err != nil {
		return nil, err
	}

	return &user,nil
}

func (r *repository) UpdateLastLogin(ctx context.Context, id uint)error{
	now := time.Now()

	return r.db.WithContext(ctx).Model(&User{}).Where("id = ?",id).Update("last_login_at",now).Error
}