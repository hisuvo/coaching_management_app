package branches

import (
	"coaching_backend/internal/domain/branches/dto"
	"coaching_backend/internal/domain/users"
)

func ToBranchResponse(branch *Branch, admin *users.User) *dto.BranchResponse {
	response := &dto.BranchResponse{
		ID:         branch.ID,
		CoachingID: branch.CoachingID,
		Name:       branch.Name,
		Code:       branch.Code,
		Phone:      branch.Phone,
		Email:      branch.Email,
		Address:    branch.Address,
		Status:     string(branch.Status),
		CreatedAt:  branch.CreatedAt,
		UpdatedAt:  branch.UpdatedAt,
	}

	if admin != nil {
		response.Admin = &dto.AdminResponse{
			ID: admin.ID,
			Name: admin.Name,
			Email: admin.Email,
			Role: string(admin.Role),
		}
	}

	return response
}

func ToBranchResponses(branches []*Branch, admins map[uint]*users.User) []*dto.BranchResponse {
	responses := make([]*dto.BranchResponse, 0, len(branches))

	for _, branch := range branches {
		admin := admins[branch.ID]
		responses = append(responses, ToBranchResponse(branch, admin))
	}

	return responses
}

func ApplyBranchUpdate(entity *Branch, req *dto.UpdateBranchRequest) {
	if req.Name != nil {
		entity.Name = *req.Name
	}

	if req.Code != nil {
		entity.Code = *req.Code
	}

	if req.Email != nil {
		entity.Email = *req.Email
	}

	if req.Address != nil {
		entity.Address = *req.Address
	}

	if req.Phone != nil {
		entity.Phone = *req.Phone
	}

	if req.Status != nil {
		entity.Status = BranchStatus(*req.Status)
	}
}