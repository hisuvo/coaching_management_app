package users

import "coaching_backend/internal/domain/users/dto"

func toUserResponse(user *User) *dto.UserResponse{
	return &dto.UserResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      string(user.Role),
		Status:    string(user.Status),
		Phone: user.Phone,
		LastLoginAt: user.LastLoginAt,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// for multiple users

func toUserResponses(users []User)[]*dto.UserResponse{
	responses := make([]*dto.UserResponse, 0, len(users))

	for i := range users {
		responses = append(responses, toUserResponse(&users[i]))
	} 

	return responses
}