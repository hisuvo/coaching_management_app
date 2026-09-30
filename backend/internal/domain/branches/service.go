package branches

import (
	"coaching_backend/internal/domain/branches/dto"
	"coaching_backend/internal/domain/users"
	"coaching_backend/internal/pkg/security"
	"context"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Service interface {
	Create(ctx context.Context, req *dto.CreateBranchRequest) (*dto.BranchResponse, error)
	FindByID(ctx context.Context, id uint) (*dto.BranchResponse, error)
	Update(ctx context.Context, id uint, req *dto.UpdateBranchRequest) (*dto.BranchResponse, error)
	Delete(ctx context.Context, id uint) (*dto.BranchResponse, error)
	FindAll(ctx context.Context) ([]*dto.BranchResponse, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(ctx context.Context, req *dto.CreateBranchRequest) (*dto.BranchResponse, error) {

	// check already have this barnch by code
	branch, err := s.repository.FindByCode(ctx, req.CoachingID, req.Code)

	if err == nil && branch != nil {
		return nil, ErrBranchAlreadyExists
	}
	
	email := strings.ToLower(strings.TrimSpace(req.Email))
	name := cases.Title(language.English).String(strings.ToLower(strings.TrimSpace(req.Name)))
	status := strings.ToUpper(req.Status)
	code := strings.ToUpper(req.Code)
	phone := req.Phone
	address := req.Address

	
	branch = &Branch{
		CoachingID: req.CoachingID,
		Name:       name,
		Code:       code,
		Phone:      phone,
		Email:      email,
		Address:    address,
		Status:  BranchStatus(status),
	}

	hashPassword, err := security.HashPassword(req.AdminPassword)

	admin := &users.User{
		CoachingID: &req.CoachingID,
		BranchID: &branch.ID,
		Name: req.AdminName,
		Email: req.AdminEmail,
		Password: hashPassword,
		Role: users.RoleBranchAdmin,
		Status: users.StatusActive,
	}

	// Save the branch
	if err := s.repository.Create(ctx, branch, admin); err != nil {
		return nil, err
	}

	// Map the branch to the response DTO
	return ToBranchResponse(branch, admin), nil

}

func (s *service) FindByID(ctx context.Context, id uint) (*dto.BranchResponse, error) {
	branch, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return ToBranchResponse(branch, nil), nil
}

func (s *service) Update(ctx context.Context, id uint, req *dto.UpdateBranchRequest) (*dto.BranchResponse, error) {
	branch, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	ApplyBranchUpdate(branch, req)

	branch, err = s.repository.Update(ctx, id, branch)
	if err != nil {
		return nil, err
	}
	return ToBranchResponse(branch, nil), nil
}

func (s *service) Delete(ctx context.Context, id uint) (*dto.BranchResponse, error) {
	branch, err := s.repository.Delete(ctx, id)

	if err != nil {
		return nil, err
	}
	
	return ToBranchResponse(branch, nil), nil
}

func (s *service) FindAll(ctx context.Context) ([]*dto.BranchResponse, error) {
	branches, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	response := ToBranchResponses(branches, nil)
	return response, nil
}
