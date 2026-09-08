package coaching

import "coaching_backend/internal/domain/coaching/dto"

func ToCoachingResponse(coaching *Coaching) *dto.CoachingResponse{
	return &dto.CoachingResponse{
		ID: coaching.ID,
		Name: coaching.Name,
		Email: coaching.Email,
		Phone: coaching.Phone,
		Domain: coaching.Domain,
		CreatedAt: coaching.CreatedAt,
		UpdatedAt: coaching.UpdatedAt,
	}
}

func ToCoachingResponses(coachings []*Coaching) []*dto.CoachingResponse {
	responses := make([]*dto.CoachingResponse,0,len(coachings))

	for i := range coachings{
		responses = append(responses, ToCoachingResponse(coachings[i]))
	}

	return responses
}