package teachers

import (
	"coaching_backend/internal/domain/teachers/dto"
	"context"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type Service interface {
	Create(ctx context.Context, coachingID uint, req *dto.CreateTeacherRequest) (*dto.TeacherResponse, error)
	GetByID(ctx context.Context, coachingID uint, id uint) (*dto.TeacherResponse, error)
	GetAll(ctx context.Context, coachingID uint) ([]*dto.TeacherResponse, error)
	Update(ctx context.Context, coachingID uint, id uint, req *dto.UpdateTeacherRequest) (*dto.TeacherResponse, error)
	Delete(ctx context.Context, coachingID uint, id uint) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

// Create creates a new teacher.
func (s *service) Create(ctx context.Context,coachingID uint,req *dto.CreateTeacherRequest) (*dto.TeacherResponse, error) {

	// Normalize employee number before checking and saving.
	employeeNo := strings.TrimSpace(strings.ToUpper(req.EmployeeNo))

	// Check whether the employee number already exists.
	existingTeacher, err := s.repository.FindByEmployeeNo(ctx, coachingID, employeeNo)

	if err != nil && !errors.Is(err, ErrTeacherNotFound) {
		return nil, err
	}

	if existingTeacher != nil {
		return nil, ErrEmployeeNoExists
	}

	// Build the Teacher entity.
	teacher := &Teacher{
		CoachingID:    coachingID,
		UserID:        req.UserID,
		BranchID:      req.BranchID,
		EmployeeNo:    employeeNo,
		Designation:   strings.TrimSpace(req.Designation),
		Qualification: strings.TrimSpace(req.Qualification),
		JoiningDate:   req.JoiningDate,
		Status:        TeacherStatus(req.Status),
	}

	// Set default status when status is not provided.
	if teacher.Status == "" {
		teacher.Status = TeacherStatusActive
	}

	// Save teacher in database.
	createdTeacher, err := s.repository.Create(ctx, teacher)

	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, ErrEmployeeNoExists
		}

		return nil, err
	}

	response := ToTeacherResponse(createdTeacher)

	return response, nil
}

// GetByID returns a teacher by ID within the coaching.
func (s *service) GetByID(ctx context.Context,coachingID uint,id uint) (*dto.TeacherResponse, error) {

	teacher, err := s.repository.GetByID(ctx, coachingID, id)

	if err != nil {
		if errors.Is(err, ErrTeacherNotFound) {
			return nil, ErrTeacherNotFound
		}

		return nil, err
	}

	// Tenant isolation check.
	if teacher.CoachingID != coachingID {
		return nil, ErrTeacherNotFound
	}

	return ToTeacherResponse(teacher), nil
}

// GetAll returns all teachers belonging to the coaching.
func (s *service) GetAll(ctx context.Context,coachingID uint) ([]*dto.TeacherResponse, error) {

	teachers, err := s.repository.GetAll(ctx, coachingID)

	if err != nil {
		return nil, err
	}

	responses := ToTeacherResponses(teachers)

	return responses, nil
}

// Update updates an existing teacher.
func (s *service) Update(ctx context.Context,coachingID uint,id uint,req *dto.UpdateTeacherRequest) (*dto.TeacherResponse, error) {

	// First get the existing teacher.
	existingTeacher, err := s.repository.GetByID(ctx,coachingID,id)

	if err != nil {
		if errors.Is(err, ErrTeacherNotFound) {
			return nil, ErrTeacherNotFound
		}

		return nil, err
	}

	// Prevent cross-coaching update.
	if existingTeacher.CoachingID != coachingID {
		return nil, ErrTeacherNotFound
	}

	// Update only the fields that were provided.
	if req.BranchID != nil {
		existingTeacher.BranchID = *req.BranchID
	}

	if req.Designation != nil {
		existingTeacher.Designation = strings.TrimSpace(*req.Designation)
	}

	if req.Qualification != nil {
		existingTeacher.Qualification = strings.TrimSpace(*req.Qualification)
	}

	if req.JoiningDate != nil {
		existingTeacher.JoiningDate = *req.JoiningDate
	}

	if req.Status != nil {
		existingTeacher.Status = TeacherStatus(*req.Status)
	}

	// Save updated teacher.
	updatedTeacher, err := s.repository.Update(ctx, coachingID, id, existingTeacher)

	if err != nil {
		return nil, err
	}

	return ToTeacherResponse(updatedTeacher), nil
}

// Delete removes a teacher.
func (s *service) Delete(ctx context.Context,coachingID uint,id uint) error {

	// Verify that the teacher exists.
	teacher, err := s.repository.GetByID(ctx, coachingID, id)

	if err != nil {
		if errors.Is(err, ErrTeacherNotFound) {
			return ErrTeacherNotFound
		}

		return err
	}

	// Prevent deleting another coaching's teacher.
	if teacher.CoachingID != coachingID {
		return ErrTeacherNotFound
	}

	return s.repository.Delete(ctx,coachingID,id)
}
