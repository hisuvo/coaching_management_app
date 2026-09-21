package auth

import (
	"coaching_backend/internal/domain/auth/dto"
	"coaching_backend/internal/domain/users"
	"context"
	"errors"
	"strings"
	"time"
)

// Task : login/refresh/logout

type Service interface {
	Login(ctx context.Context, req *dto.LoginRequest, userAgent string, ipAddress string,) (*dto.LoginResult, error)
	Refresh(ctx context.Context,refreshToken string,userAgent string,ipAddress string,) (*dto.RefreshResult, error)
	Logout(ctx context.Context,refreshToken string,) error
	LogoutAll(ctx context.Context,userID uint,) error
}

type service struct {
	userRepository users.Repository
	authRepository Repository
	tokenManager *TokenManager

}

func NewService(userRepository users.Repository,authRepository Repository,tokenManager *TokenManager,) Service {
	return &service{
		userRepository: userRepository,
		authRepository: authRepository,
		tokenManager:   tokenManager,

	}
}

func (s *service) Login(ctx context.Context, req *dto.LoginRequest,userAgent string, ipAddress string,) (*dto.LoginResult, error) {
	email := strings.ToLower(strings.TrimSpace(req.Email))

	user, err := s.userRepository.FindByEmail(ctx,email)


	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if user.Status == "" {
		user.Status = users.StatusActive
	}

	if user.Status != users.StatusActive {
		return nil, errors.New("user account is not active")
	}

	if err := ComparePassword(user.PasswordHash, req.Password); err != nil {
		return nil, errors.New("Password does NOT match")
	}

	refreshToken, err := GenerateRefreshToken()

	if err != nil {
		return nil, err
	}

	refreshTokenHash := HashRefreshToken(refreshToken)

	session := &AuthSession{
		UserID:            user.ID,
		CoachingID:        user.CoachingID,
		RefreshTokenHash:  refreshTokenHash,
		UserAgent:         userAgent,
		IPAddress:         ipAddress,
		ExpiresAt:         time.Now().Add(
			s.tokenManager.RefreshExpires,
		),
	}

	if err := s.authRepository.CreateSession(ctx,session); err != nil {
		return nil, err
	}

	accessToken, expiresIn, err := s.tokenManager.GenerateAccessToken(
		user.ID,
		user.CoachingID,
		string(user.Role),
		session.ID,
	)

	if err != nil {
		return nil, err
	}

	if err := s.userRepository.UpdateLastLogin(
		ctx,
		user.ID,
	); err != nil {
		return nil, err
	}

	return &dto.LoginResult{
		AccessToken:  accessToken,
		ExpiresIn:    expiresIn,
		RefreshToken: refreshToken,
		UserID:       user.ID,
		CoachingID:   user.CoachingID,
		Role:         string(user.Role),
		Name:         user.Name,
		Email:        user.Email,
	}, nil
}

func (s *service) Refresh( ctx context.Context, refreshToken string, userAgent string, ipAddress string) (*dto.RefreshResult, error) {

	if refreshToken == "" {
		return nil, errors.New("refresh token is required")
	}

	tokenHash := HashRefreshToken(refreshToken)

	session, err :=
		s.authRepository.GetSessionByTokenHash(
			ctx,
			tokenHash,
		)

	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	if session.RevokedAt != nil {
		return nil, errors.New("refresh session has been revoked")
	}

	if time.Now().After(session.ExpiresAt) {
		return nil, errors.New("refresh token has expired")
	}

	user, err := s.userRepository.GetByID(
		ctx,
		session.UserID,
	)

	if err != nil {
		return nil, errors.New("user not found")
	}

	if user.Status != users.StatusActive {
		return nil, errors.New("user account is not active")
	}

	newRefreshToken, err := GenerateRefreshToken()

	if err != nil {
		return nil, err
	}

	newHash := HashRefreshToken(newRefreshToken)

	newExpiresAt := time.Now().Add(
		s.tokenManager.RefreshExpires,
	)

	if err := s.authRepository.UpdateRefreshToken(
		ctx,
		session.ID,
		newHash,
		newExpiresAt,
	); err != nil {
		return nil, err
	}

	accessToken, expiresIn, err :=
		s.tokenManager.GenerateAccessToken(
			user.ID,
			user.CoachingID,
			string(user.Role),
			session.ID,
		)

	if err != nil {
		return nil, err
	}

	return &dto.RefreshResult{
		AccessToken:  accessToken,
		ExpiresIn:    expiresIn,
		RefreshToken: newRefreshToken,
	}, nil
}

func (s *service) Logout(ctx context.Context, refreshToken string,) error {

	if refreshToken == "" {
		return nil
	}

	tokenHash := HashRefreshToken(refreshToken)

	session, err :=
		s.authRepository.GetSessionByTokenHash(
			ctx,
			tokenHash,
		)

	if err != nil {
		return nil
	}

	return s.authRepository.RevokeSession(
		ctx,
		session.ID,
	)
}

func (s *service) LogoutAll(ctx context.Context, userID uint,) error {
	return s.authRepository.RevokeAllUserSession(
		ctx,userID,
	)
}

