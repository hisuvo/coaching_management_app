package submissions

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoute(e *echo.Echo, db *gorm.DB) {

	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)

	group := e.Group("/api/v1/submissions")

	group.POST("",handler.Create)
}