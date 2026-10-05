package coachingsubject

import (
	"coaching_backend/internal/domain/coaching"
	"coaching_backend/internal/domain/coachingSubject/dto"
	"coaching_backend/internal/domain/subjects"
	"context"
	"errors"
)

type Service interface {
	Create(ctx context.Context, req *dto.CreateCoachingSubjectRequest) (*dto.CreateCoachingSubjectResponse, error)
	GetSingleById(ctx context.Context, id uint) (*dto.CoachingSubjectResponse, error)
	GetOwnCoachingSubject(ctx context.Context, coaching_id *uint) ([]*dto.CoachingSubjectResponse, error)
	GetAll(ctx context.Context)([]*dto.CoachingSubjectResponse, error)
	Update(ctx context.Context, id uint, status *string) (*dto.UpdateCoachingSubjectRequest, error)
	Delete(ctx context.Context, id uint) error
}

type service struct {
	repository Repository
	subjectRepo subjects.Repoistory
	coachingRepo coaching.Repository
}

func NewService(repository Repository, subjectRepo subjects.Repoistory, coachingRepo coaching.Repository) Service {
	return &service{
		repository: repository,
		subjectRepo: subjectRepo,
		coachingRepo: coachingRepo,
	}
}

func (s *service) Create(ctx context.Context, req *dto.CreateCoachingSubjectRequest) (*dto.CreateCoachingSubjectResponse, error){
	
	var subject *subjects.Subject

	// Check subject is exists in database
	existsSubject, err := s.subjectRepo.FindByCode(ctx, req.Code)

	if err == nil {
		subject = existsSubject
	} else if errors.Is(err, subjects.ErrSubjectNotFound){
		subject = &subjects.Subject{
			Name: req.Name,
			Code: req.Code,
			Description: req.Description,
		}

		// subject created here
		if err := s.subjectRepo.Create(ctx, subject); err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}
	
	var coaching_subjects *CoachingSubject

	existsCoachingSubject, err := s.repository.FindCoachingSubjectByCoadeAndId(ctx, req.CoachingID, subject.ID)

	if err == nil {
		coaching_subjects = existsCoachingSubject
		return nil, ErrCoachingSubjectExist
	} else if errors.Is(err, ErrCoachingSubjectNotFound) {
		coaching_subjects = &CoachingSubject{
			CoachingID: req.CoachingID,
			SubjectID: subject.ID,
			Status: string(CoachingSubjectStatusActive),
			Subject: *subject,
		}

		if err := s.repository.Create(ctx, coaching_subjects); err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}

	response := ToCreateCoachingSubjectResponse(coaching_subjects)

	return response, nil
}

func (s *service) GetOwnCoachingSubject(ctx context.Context, coaching_id *uint) ([]*dto.CoachingSubjectResponse, error) {
	result, err := s.repository.GetOwnCoachingSubject(ctx, coaching_id)

	if err != nil{
		return nil, err
	}

	response := ToCoachingSubjectResponses(result)

	return response, nil
}

func (s *service) GetSingleById(ctx context.Context, id uint) (*dto.CoachingSubjectResponse, error){
	
	coaching_subject, err := s.repository.GetSingleById(ctx, id)

	if err != nil {
		return nil, err
	}

	subject, err := s.subjectRepo.GetById(ctx, coaching_subject.SubjectID)

	if err != nil {
		return nil, err
	}

	response := &dto.CoachingSubjectResponse{
		ID: coaching_subject.ID,
		Status: coaching_subject.Status,
		Subject: subjects.ToSubjectResponse(subject),
	}
	return response,nil
}

func (s *service) GetAll(ctx context.Context)([]*dto.CoachingSubjectResponse, error) {
	coaching_subject, err := s.repository.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	reponse := ToCoachingSubjectResponses(coaching_subject)

	return reponse, nil
}

func (s *service) Update(ctx context.Context, id uint, status *string) (*dto.UpdateCoachingSubjectRequest, error){
	updateStatus, err := s.repository.Upate(ctx, id, status)

	if err != nil {
		return nil, err
	}

	response := &dto.UpdateCoachingSubjectRequest{
		Status: &updateStatus.Status,
	}

	return response, nil
}

func (s *service) Delete(ctx context.Context, id uint) error {
	
	exists, err := s.repository.GetSingleById(ctx, id)

	if err != nil {
		return err
	}

	if err := s.repository.DeleteCoaching(ctx, exists.ID); err != nil {
		return err
	}

	return nil
}