package auth

import "coaching_backend/internal/domain/auth/dto"

func ToLoginResopnse(user *dto.LoginResponse) *dto.LoginResponse {
	return &dto.LoginResponse{
		AccessToken: user.AccessToken,
		TokenType: user.TokenType,
		ExpiresIn: user.ExpiresIn,
		User: dto.UserResponse{
			ID: user.User.ID,
			CoachingID: user.User.CoachingID,
			Name: user.User.Name,
			Email: user.User.Email,
			Role: user.User.Role,
			Status: user.User.Status,
			Phone: user.User.Phone,
		},
	}
}