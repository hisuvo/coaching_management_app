package batches

import "coaching_backend/internal/domain/batches/dto"

func ToBatchResponse(batch *Batch) *dto.BatchResponse {
	return &dto.BatchResponse{
	}
}

func ToBatchResponses (batches []*Batch) []*dto.BatchResponse {
	response := make([]*dto.BatchResponse, 0, len(batches))

	for i := range batches {
		return append(response,ToBatchResponse(batches[i]))
	}

	return response
}