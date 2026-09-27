package branches

import "coaching_backend/internal/domain/branches/dto"

func ToBranchResponse(branch *Branch) *dto.BranchResponse {
	return &dto.BranchResponse{
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
}

func ToBranchResponses(branches []*Branch) []*dto.BranchResponse {
	responses := make([]*dto.BranchResponse, 0, len(branches))

	for _, branch := range branches {
		responses = append(responses, ToBranchResponse(branch))
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