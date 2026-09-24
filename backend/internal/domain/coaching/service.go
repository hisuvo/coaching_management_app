package coaching

import (
	"coaching_backend/internal/domain/coaching/dto"
	"context"
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	Create(ctx context.Context,req *dto.CreateCoachingRequest) (*dto.CoachingResponse, error)
	GetById(ctx context.Context,id uint) (*dto.CoachingResponse, error)
	GetAll(ctx context.Context) ([]*dto.CoachingResponse, error)
	Update(ctx context.Context, id uint, req *dto.UpdateCoachingRequest) (*dto.CoachingResponse, error)
	Delete(ctx context.Context, id uint) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service{
	return &service{
		repository: repository,
	}
}

func (s *service) Create(ctx context.Context,req *dto.CreateCoachingRequest) (*dto.CoachingResponse, error) {
	// Check whether a coaching already exists with this email.
	_, err := s.repository.FindByEmail(ctx, req.Email)

	// No error means the email already exists.
	if err == nil {
		return nil, ErrCoachingEmailExist
	}

	// Continue only when the coaching was actually not found.
	if !errors.Is(err, ErrCoachingNotFound) {
		return nil, err
	}

	// Generate a URL-friendly slug from the coaching name.
	coachingSlug := slug.Make(req.Name)

	// Generate the platform subdomain.
	domain := fmt.Sprintf("%s.coachinghub.com",coachingSlug)

	// Build the coaching entity.
	coaching := &Coaching{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Domain:   domain,
		Slug:     coachingSlug,
		LogoURL: req.LogoURL,
		Address: req.Address,
		TimeZone: req.TimeZone,
	}

	fmt.Print()

	// Create the coaching.
	if err := s.repository.Create(ctx, coaching); err != nil {
		return nil, err
	}

	// Convert entity to response DTO.
	response := ToCoachingResponse(coaching)

	return response, nil
}

func (s *service) GetById(ctx context.Context,id uint) (*dto.CoachingResponse, error) {
	caching, err := s.repository.GetById(ctx,id)

	if err != nil {
		return nil, err
	}

	response := ToCoachingResponse(caching)

	return response, nil
}

func (s *service) GetAll(ctx context.Context,) ([]*dto.CoachingResponse, error) {
	cachings, err := s.repository.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	responses := ToCoachingResponses(cachings)

	return responses, nil
}

func (s *service)  Update(ctx context.Context, id uint, req *dto.UpdateCoachingRequest) (*dto.CoachingResponse, error){
	existing, err := s.repository.GetById(ctx, id)
	
	if err != nil {
		return nil, err
	}
	
	// Update only field provide by client
	ApplyCoachingUpdate(existing, req)

	// Save updated coaching.
	updated, err := s.repository.Update(ctx,id, existing)

	if err != nil {
		return nil, err
	}

	response := ToCoachingResponse(updated)

	return response,nil
}

func (s *service) Delete(ctx context.Context, id uint) error {
	if err := s.repository.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}