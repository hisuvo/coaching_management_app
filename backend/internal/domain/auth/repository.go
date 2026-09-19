package auth

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

// Task : session DB operations

type Repository interface {
	CreateSession(ctx context.Context, session *AuthSession) error
	GetSessionByTokenHash(ctx context.Context, tokenHash string) (*AuthSession, error)
	RevokeSession(ctx context.Context, sessionId uint) error
	RevokeAllUserSession(ctx context.Context, userId uint) error
	UpdateRefreshToken(ctx context.Context, sessionID uint, tokenHash string, expiresAt time.Time) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateSession (ctx context.Context, session *AuthSession) error {
	return r.db.WithContext(ctx).Create(session).Error
}

func (r *repository) GetSessionByTokenHash(ctx context.Context, tokenHash string) (*AuthSession, error){
	var session AuthSession

	err := r.db.WithContext(ctx).Where("refresh_token_hash = ?", tokenHash).First(&session).Error

	if errors.Is(err,gorm.ErrRecordNotFound){
		return nil, gorm.ErrRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (r *repository) RevokeSession(ctx context.Context, sessionId uint) error {
	now := time.Now()
	return r.db.WithContext(ctx).Model(&AuthSession{}).Where("session_id = ?", sessionId).Update("revoke_at", now).Error
}

func (r *repository) RevokeAllUserSession(ctx context.Context, userId uint) error {
	now := time.Now()
	return r.db.WithContext(ctx).Model(&AuthSession{}).Where("user_id = ? AND revoked_at IS NULL", userId).Update("revoked_at", now).Error
}

func (r *repository) UpdateRefreshToken(ctx context.Context, sessionID uint, tokenHash string, expiresAt time.Time) error{
	return r.db.WithContext(ctx).
		Model(&AuthSession{}).
		Where("id = ?", sessionID).
		Updates(map[string]interface{}{
			"refresh_token_hash": tokenHash,
			"expires_at":         expiresAt,
		}).
		Error
}
