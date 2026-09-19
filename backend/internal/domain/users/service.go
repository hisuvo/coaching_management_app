package users

import (
	"coaching_backend/internal/domain/users/dto"
	"context"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	Register(ctx context.Context, req dto.CreateUserRequest) (*dto.UserResponse, error)
	FindByEmail(ctx context.Context, email string)(*dto.UserResponse, error)
	GetByID(ctx context.Context, id uint) (*dto.UserResponse, error)
}

type service struct {
	repository Repository
}

// NewService is a constructor function
func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Register(ctx context.Context, req dto.CreateUserRequest)(*dto.UserResponse, error){

	// Normalizw input
	name := strings.TrimSpace(req.Name)
	email := strings.ToLower(req.Email)
	phone := strings.TrimSpace(req.Phone)

	// check duplicate email
	existingUser, err := s.repository.FindByEmail(ctx,email)

	if err != nil && !errors.Is(err,gorm.ErrRecordNotFound) {
		return nil, err
	}

	if existingUser != nil {
		return nil, ErrDuplicateEmail
	}

	// Hash Password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.PasswordHash),bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	user := &User{
		Name: name,
		Email: email,
		PasswordHash: string(passwordHash),
		Phone: phone,
		LastLoginAt: nil,
	}


	if err := s.repository.Create(ctx, user); err != nil {
		return nil, err
	}

	response := toUserResponse(user)

	return response, nil
}

func (s *service) FindByEmail(ctx context.Context, email string) (*dto.UserResponse, error){
	user, err := s.repository.FindByEmail(ctx, email)

	if err != nil {
		if errors.Is(err, ErrEmailNotFound) {
			return nil, ErrEmailNotFound
		}
		return nil, err
	}

	return toUserResponse(user), nil
}

func (s *service) GetByID(ctx context.Context, id uint) (*dto.UserResponse, error){
	user, err := s.repository.GetByID(ctx, id)

	if err != nil {
		if errors.Is(err, ErrUserNotRound) {
			return nil, ErrUserNotRound
		}
		return nil, err
	}

	return toUserResponse(user), nil
}