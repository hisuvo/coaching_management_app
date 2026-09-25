package batches

import "coaching_backend/internal/domain/batches/dto"

func ToBatchResponse(batch *dto.CreateBatchRequest) *dto.BatchResponse {
	return &dto.BatchResponse{
		Name: batch.Name,
		Email: batch.Email,
		Phone: batch.Phone,
		Code: batch.Code,
		Address: batch.Address,
		City: batch.City,
		Division: batch.Division,
		Status: string(batch.Status),
		TimeZone: batch.TimeZone,
		OpeningTime: batch.OpeningTime,
		ClosingTime: batch.ClosingTime,
	}
}

func ToBatchResponses (batches []*dto.CreateBatchRequest) []*dto.BatchResponse {
	response := make([]*dto.BatchResponse, 0, len(batches))

	for i := range batches {
		return append(response,ToBatchResponse(batches[i]))
	}

	return response
}