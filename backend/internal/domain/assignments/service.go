package assignments

import (
	"coaching_backend/internal/domain/assignments/dto"
	"errors"
	"time"
)

type Service interface {
	GetAll() ([]*dto.AssignmentResponse, error)
	Create(req *dto.CreateAssignmentRequest, createdBy uint) (*dto.AssignmentResponse, error)
	GetById(assignmentId string) (*dto.AssignmentResponse, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service{
	return &service{
		repository: repository,
	}
}

func (s *service) GetAll() ([]*dto.AssignmentResponse, error) {
	assignments, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}

	return ToAssignmentResponses(assignments), nil
}

func (s *service) Create(req *dto.CreateAssignmentRequest, createdBy uint) (*dto.AssignmentResponse, error){

	assignment := &Assignment{
		Title: req.Title,
		Description: req.Description,
		SubjectID:   req.SubjectID,
		CoachingID:  req.CoachingID,
		CreatedBy:   createdBy,
		Status:      AssignmentStatusDraft,	
	}

	if req.Status != "" {
		assignment.Status = AssignmentStatus(req.Status)
	}

	// Todo: here time zone add later
	if req.DueDate != nil {
		dueDate, err := time.Parse(time.DateOnly, *req.DueDate)

		if err != nil {
			return nil , err
		}

		assignment.DueDate = &dueDate
	}

	if err := s.repository.Create(assignment); err != nil {
		return nil, err
	}

	return ToAssignmentResponse(assignment), nil
}

func (s *service) GetById(assignmentId string) (*dto.AssignmentResponse, error) {
	if assignmentId == "" {
		return nil, errors.New("assignment id need for signle assignment")
	}

	assignment, err := s.repository.GetById(assignmentId)

	if err != nil {
		return nil, err
	}

	response := ToAssignmentResponse(assignment)

	return response, nil
}