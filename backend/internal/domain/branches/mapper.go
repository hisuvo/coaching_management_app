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
		City:       branch.City,
		Country:    branch.Country,
		Status:     branch.Status,
		CreatedAt:  branch.CreatedAt,
		UpdatedAt:  branch.UpdatedAt,
	}
}

func ToBranchResponses(branches []*Branch) []*dto.BranchResponse {

	response := make([]*dto.BranchResponse, 0, len(branches))

	for i := range branches{
		return append(response, ToBranchResponse(branches[i]))
	}

	return response
}