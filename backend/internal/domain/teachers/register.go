package teachers

import (
	"coaching_backend/internal/domain/auth"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoute(api *echo.Group, db *gorm.DB, authMiddleware echo.MiddlewareFunc){
	repository := NewRegister(db)
	service := NewService(repository)
	handler := NewHandler(service)

	api.POST("/teacher",handler.Create, authMiddleware, auth.RequireRoles("SUPER_ADMIN"))
	api.GET("/teachers", handler.GetAll, authMiddleware)
}