package coaching

import "coaching_backend/internal/domain/coaching/dto"

func ToCoachingResponse(coaching *Coaching) *dto.CoachingResponse{
	return &dto.CoachingResponse{
		ID: coaching.ID,
		Name: coaching.Name,
		Email: coaching.Email,
		Phone: coaching.Phone,
		Domain: coaching.Domain,
		Slug: coaching.Slug,
		LogoURL: coaching.LogoURL,
		Address: coaching.Address,
		TimeZone: coaching.TimeZone,
		Status: string(coaching.Status),
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

func ApplyCoachingUpdate(entity *Coaching, req *dto.UpdateCoachingRequest) {
	if req.Name != nil {
		entity.Name = *req.Name
	}
	
	if req.Email != nil {
		entity.Email = *req.Email
	}

	if req.Phone != nil {
		entity.Phone = *req.Phone
	}

	if req.LogoURL != nil {
		entity.LogoURL = *req.LogoURL
	}

	if req.Status != nil {
		entity.Status = CoachingStatus(*req.Status)
	}

	if req.TimeZone != nil {
		entity.TimeZone = *req.TimeZone
	}
}