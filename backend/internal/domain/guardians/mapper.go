package guardians

import "coaching_backend/internal/domain/guardians/dto"

func ToGuardianResponse(guardian *Guardian) *dto.GuardianResponse {
	return &dto.GuardianResponse{}
}

func ToGuardianResponses(guardians []*Guardian) []*dto.GuardianResponse {
	response := make([]*dto.GuardianResponse, 0, len(guardians))

	for i := range guardians {
		return append(response, ToGuardianResponse(guardians[i]))
	}

	return response
}