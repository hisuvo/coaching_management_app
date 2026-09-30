package coachingsubject

import (
	"coaching_backend/internal/domain/coachingSubject/dto"
	"context"
)

type Service interface {
	Create(ctx context.Context, req *dto.CreateCoachingSubjectRequest) (*dto.CoachingSubjectResponse, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(ctx context.Context, req *dto.CreateCoachingSubjectRequest) (*dto.CoachingSubjectResponse, error){
	var coaching_subjects dto.CoachingSubjectResponse

	coaching_subjects = dto.CoachingSubjectResponse{
		CoachingID: req.CoachingID,
		SubjectID: req.SubjectID,
		Status: string(CoachingSubjectStatusActive),
	}

	return &coaching_subjects, nil
}