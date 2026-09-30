package coaching

import (
	"coaching_backend/internal/domain/coaching/dto"
	"coaching_backend/internal/domain/users"
)

func ToCoachingResponse(coaching *Coaching, admin *users.User) *dto.CoachingResponse{
	response := &dto.CoachingResponse{
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

func ToCoachingResponses(coachings []*Coaching, admins map[uint]*users.User) []*dto.CoachingResponse {
	responses := make([]*dto.CoachingResponse,0,len(coachings))

	for _, coaching := range coachings{
		admin := admins[coaching.ID]
		responses = append(responses, ToCoachingResponse(coaching, admin))
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