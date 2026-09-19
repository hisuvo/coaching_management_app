package assignments

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoute(e *echo.Echo, db *gorm.DB) {

	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)

	route := e.Group("/api/v1/assignments")

	route.POST("", handler.Create)
	route.GET("", handler.GetAll)
	route.GET("/:assignmentId", handler.GetById)
}