package subjects

import (
	"coaching_backend/internal/domain/subjects/dto"
	"context"
	"errors"
)

type Service interface {
	CreateSubject(ctx context.Context, req *dto.CreateSubjectRequest) (*dto.SubjectResponse, error)
	GetAll()([]*dto.SubjectResponse, error)
	GetById(ctx context.Context, subjectId uint)(*dto.SubjectResponse, error)
	Update(subjectId string, req *dto.UpdateSubjectRequest)(*dto.SubjectResponse, error)
	Delete(subjectId string) (*dto.SubjectResponse, error)
}

type service struct{
	// db  *gorm.DB
	repository Repoistory
}

func NewService(repository Repoistory) Service{
	return &service{
		repository: repository,
	}
}

func (s *service) CreateSubject(ctx context.Context, req *dto.CreateSubjectRequest) (*dto.SubjectResponse, error){
	var subject *Subject
	
	// Check subject already exists in database
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

func (s *service) GetById(ctx context.Context, subjectId uint)(*dto.SubjectResponse, error) {
	subject, err := s.repository.GetById(ctx, subjectId)

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