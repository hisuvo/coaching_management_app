package branches

import (
	"coaching_backend/internal/domain/branches/dto"
	"context"
	"strings"
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
	name := strings.ToUpper(req.Name)
	status := strings.ToUpper(req.Status)
	code := strings.ToUpper(req.Code)
	country := strings.ToUpper(req.Country)
	phone := req.Phone
	address := req.Address
	city := req.City
	
	branch = &Branch{
		CoachingID: req.CoachingID,
		Name:       name,
		Code:       code,
		Phone:      phone,
		Email:      email,
		Address:    address,
		City:       city,
		Country: country,
		Status:  BranchStatus(status),
	}

	// Save the branch
	newBranch, err := s.repository.Create(ctx, branch)
	if err != nil {
		return nil, err
	}

	// Map the branch to the response DTO
	return ToBranchResponse(newBranch), nil

}

func (s *service) FindByID(ctx context.Context, id uint) (*dto.BranchResponse, error) {
	branch, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return ToBranchResponse(branch), nil
}

func (s *service) Update(ctx context.Context, id uint, req *dto.UpdateBranchRequest) (*dto.BranchResponse, error) {
	branch, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		branch.Name = *req.Name
	}

	if req.Code != nil {
		branch.Code = *req.Code
	}

	if req.Phone != nil {
		branch.Phone = *req.Phone
	}

	if req.Email != nil {
		branch.Email = *req.Email
	}

	if req.Address != nil {
		branch.Address = *req.Address
	}

	if req.City != nil {
		branch.City = *req.City
	}
	if req.Country != nil {
		branch.Country = *req.Country
	}
	if req.Status != nil {
		branch.Status = BranchStatus(*req.Status)
	}

	branch, err = s.repository.Update(ctx, id, branch)
	if err != nil {
		return nil, err
	}
	return ToBranchResponse(branch), nil
}

func (s *service) Delete(ctx context.Context, id uint) (*dto.BranchResponse, error) {
	branch, err := s.repository.Delete(ctx, id)
	if err != nil {
		return nil, err
	}
	return ToBranchResponse(branch), nil
}

func (s *service) FindAll(ctx context.Context) ([]*dto.BranchResponse, error) {
	branches, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	response := ToBranchResponses(branches)
	return response, nil
}