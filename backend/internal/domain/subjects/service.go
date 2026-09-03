package subjects

import "coaching_backend/internal/domain/subjects/dto"

type Service interface {
	CreateSubject(req *dto.CreateSubjectRequest) (*dto.SubjectResponse, error)
}

type service struct{
	repository Repoistory
}

func NewService(repository Repoistory) Service{
	return &service{
		repository: repository,
	}
}

func (s *service) CreateSubject(req *dto.CreateSubjectRequest) (*dto.SubjectResponse, error){
	subject := Subject{
		Name: req.Name,
		Code: req.Code,
		Description: req.Description,
		Status: SubjectStatus(req.Status),
	}
	err := s.repository.Create(&subject)
	
	if err != nil {
		return nil, err
	}
	return ToSubjectResponse(&subject), nil
}