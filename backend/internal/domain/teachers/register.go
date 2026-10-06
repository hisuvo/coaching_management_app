package teachers

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoute(api *echo.Group, db *gorm.DB, authMiddleware echo.MiddlewareFunc){
	repository := NewRegister(db)
	service := NewService(repository)
	handler := NewHandler(service)

	api.POST("/teacher",handler.Create, authMiddleware)
}