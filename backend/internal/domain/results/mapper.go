package results

import "coaching_backend/internal/domain/results/dto"

func ToResultResponse(result *Result) *dto.ResultResponse {
	return &dto.ResultResponse{}
}

func ToResultResponses (results []*Result) []*dto.ResultResponse {
	response := make([]*dto.ResultResponse, 0, len(results))

	for i := range results {
		return append(response, ToResultResponse(results[i]))
	}
	
	return response
}