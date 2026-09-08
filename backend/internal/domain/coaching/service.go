package coaching

import (
	"coaching_backend/internal/apperror"
	"coaching_backend/internal/domain/coaching/dto"
	"errors"
)

type Service interface {
	Create(req *dto.CreateCoachingRequest) (*dto.CoachingResponse, error)
	GetById(id string) (*dto.CoachingResponse, error)
	GetAll() ([]*dto.CoachingResponse, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service{
	return &service{
		repository: repository,
	}
}

func (s *service) Create(req *dto.CreateCoachingRequest) (*dto.CoachingResponse, error){
	coaching := &Coaching{
		Name: req.Name,
		Email: req.Email,
		Phone: req.Phone,
		Domain: req.Domain,
	}

	if err := s.repository.Create(coaching); err != nil {
		return nil, err
	}

	response := ToCoachingResponse(coaching)

	return response, nil
}

func (s *service) GetById(id string) (*dto.CoachingResponse, error) {
	caching, err := s.repository.GetById(id)

	if err != nil {
		if errors.Is(ErrNotFoundCoaching, err){
			return nil, apperror.NotFound("Coaching not found")
		}

		return nil, err
	}

	response := ToCoachingResponse(caching)

	return response, nil
}

func (s *service) GetAll() ([]*dto.CoachingResponse, error) {
	cachings, err := s.repository.GetAll()

	if err != nil {
		if errors.Is(ErrNotFoundCoaching, err){
			return nil, apperror.NotFound("Coaching not found")
		}

		return nil, err
	}

	responses := ToCoachingResponses(cachings)

	return responses, nil
}
