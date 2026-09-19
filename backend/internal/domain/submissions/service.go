package submissions

import (
	"coaching_backend/internal/domain/submissions/dto"
	"errors"
	"strconv"
)

type Service interface {
	Create(coachingId string,req *dto.CreateSubmissionRequest) (*dto.SubmissionResponse, error)
}

type service struct {
	repository Repository
}

func NewService (repository Repository) Service {
	return &service{
		repository: repository,
	}
}

// Create creates a submission.
func (s *service) Create(coachingId string, req *dto.CreateSubmissionRequest) (*dto.SubmissionResponse, error) {
	
	// validate request pointer
	if req == nil {
		return nil, errors.New("Submission request can not be nil")
	}

	// A submission must contain either a file or text
	if *req.FileURL == "" && (req.TextAnswer == nil || *req.TextAnswer == "") {
		return nil, errors.New("submission must containe file url and text answer")
	}

	// Check whether this student already submitted this assignment
	existing, err := s.repository.FindByStudentAndAssignment(
		strconv.Itoa(int(req.StudentID)), 
		strconv.Itoa(int(req.AssignmentID)), 
		coachingId)

	// if an existing recoard is found, prevent duplicate submission
	if err == nil && existing != nil {
		return nil, errors.New("student has already submitted this assignment")
	}

	// Ignore not-found error because no existing submission is expected.
	if err != nil && IsNotFound(err) {
		return nil, err
	}

	// defult status 
	status := SubmissionStatusSubmitted

	submission := &Submission{
		AssignmentID: req.AssignmentID,
		StudentID: req.StudentID,
		TextAnswer: req.TextAnswer,
		FileURL: *req.FileURL,
		Status: string(status),
	}
	
	response := ToSubmissionResponse(submission)
	return response, nil
}