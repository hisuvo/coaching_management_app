package users

import (
	"coaching_backend/internal/apperror"
	"coaching_backend/internal/domain/users/dto"
	"errors"

	"gorm.io/gorm"
)

type Service interface {
	Register(req dto.CreateUserRequest) (*dto.UserResponse, error)
	FindByEmail(email string)(*dto.UserResponse, error)
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

func (s *service) Register(req dto.CreateUserRequest)(*dto.UserResponse, error){

	user := &User{
		Name: req.Name,
		Email: req.Email,
		Password: req.Password,
		Role: Role(req.Role),
	}

	if err := s.repository.Create(user); err != nil {
		return nil, err
	}

	response := toUserResponse(user)

	return response, nil
}

func (s *service) FindByEmail(email string) (*dto.UserResponse, error){
	user, err := s.repository.FindByEmail(email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.NotFound("User not found")
		}
		return nil, err
	}

	return toUserResponse(user), nil
}