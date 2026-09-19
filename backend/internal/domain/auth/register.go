package auth

import (
	"coaching_backend/internal/config"
	"coaching_backend/internal/domain/users"
	"time"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, config *config.Config) {
	userRepo := users.NewRepository(db)
	authRepo := NewRepository(db)
	tokenManager := NewTokenManager(
		config.JWT_ACCESS_SECRET,
		"coaching-management-api",
		15*time.Minute,
		7*24*time.Hour,
	)
	authService := NewService(userRepo, authRepo, tokenManager)
	handler := NewHandler(authService, *config)

	auth := e.Group("/api/v1/auth")

	auth.POST("/login", handler.Login)

	// auth.POST(
	// 	"/refresh",
	// 	handler.Refresh,
	// )

	// auth.POST(
	// 	"/logout",
	// 	handler.Logout,
	// )
}