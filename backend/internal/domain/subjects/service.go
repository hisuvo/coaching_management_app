package subjects

import (
	coachingsubject "coaching_backend/internal/domain/coachingSubject"
	"coaching_backend/internal/domain/subjects/dto"
	"context"
	"errors"
)

type Service interface {
	CreateSubject(ctx context.Context, req *dto.CreateSubjectRequest) (*dto.SubjectResponse, error)
	GetAll()([]*dto.SubjectResponse, error)
	GetById (subjectId string) (*dto.SubjectResponse, error)
	Update(subjectId string, req *dto.UpdateSubjectRequest) (*dto.SubjectResponse, error)
	Delete(subjectId string) (*dto.SubjectResponse, error)
}

type service struct{
	// db  *gorm.DB
	repository Repoistory
	coaching_subject_repository coachingsubject.Repository
}

func NewService(repository Repoistory,coaching_subject_repository coachingsubject.Repository) Service{
	return &service{
		repository: repository,
		coaching_subject_repository: coaching_subject_repository,
	}
}

func (s *service) CreateSubject(ctx context.Context, req *dto.CreateSubjectRequest) (*dto.SubjectResponse, error){
	var subject *Subject
	
	// If subject does not exists, create it
	existingSubject, err := s.repository.FindByCode(ctx, req.Code)

	if err == nil {
		// subject already exists
		// we don't create another subject
		subject = existingSubject
	} else if errors.Is(err, ErrSubjectNotFound) {
		// subject dosen't exist
		// create a new subject
		subject = &Subject{
			Name: req.Name,
			Code: req.Code,
			Description: req.Description,
		}

		if err := s.repository.Create(ctx, subject); err != nil {
			return nil , err
		}
	} else {
		return nil, err
	}

	// Subject now definitely exists.
	// Create the relationship between Coaching and Subject.
	coachingSubject := &coachingsubject.CoachingSubject{
		CoachingID: req.CoachingID,
		SubjectID:  subject.ID,
		Status:     "active",
	}

	if err := s.coaching_subject_repository.Create(ctx, coachingSubject); err != nil {
		return nil, err
	}

	return ToSubjectResponse(subject), nil
}

func (s *service) GetAll() ([]*dto.SubjectResponse, error) {
	subjects, err := s.repository.GetAll()

	if err != nil {
		return nil, ErrSubjectNotFound
	}

	response := ToSubjectResponses(subjects)

	return response, nil
}

func (s *service) GetById(subjectId string) (*dto.SubjectResponse, error) {
	subject, err := s.repository.GetById(subjectId)

	if err != nil {
		return nil, err
	}

	response := ToSubjectResponse(subject)
	return response, nil
}

func (s *service) Update(subjectId string, req *dto.UpdateSubjectRequest) (*dto.SubjectResponse, error){
	subject := &Subject{
		Name: *req.Name,
		Code: *req.Code,
		Description: req.Description,
	}

	updateRes, err := s.repository.Update(subjectId, subject)

	if err !=nil {
		return nil, err
	}

	response := ToSubjectResponse(updateRes)

	return response, nil
}

func (s *service) Delete(subjectId string) (*dto.SubjectResponse, error) {
	subject, err := s.repository.Delete(subjectId)

	if err != nil {
		return nil, err
	}

	response := ToSubjectResponse(subject)
	return response, nil
}