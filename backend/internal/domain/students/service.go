package students

import (
	"coaching_backend/internal/domain/students/dto"
)

type Service interface {
	Create(req *dto.CreateStudentRequest) (*dto.StudentResponse, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service{
	return &service{
		repository: repository,
	}
}

func (s *service) Create(req *dto.CreateStudentRequest) (*dto.StudentResponse, error){
	student := &Student{
		CoachingID: req.CoachingID,
		Name: req.Name,
		Class:req.Class,		
		Session: req.Session,
		Email: req.Email,
		Phone: req.Phone,
		SchoolName: req.SchoolName,
		Address: req.Address,
	}

	if err := s.repository.Create(student); err != nil {
		return nil, err
	}
	

	response := ToStudentResponse(student)

	return response, nil
}