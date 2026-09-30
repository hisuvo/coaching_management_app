package coachingsubject

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoute(api *echo.Group, db *gorm.DB) {
	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)

	api.POST("/coaching_subjects", handler.Create)
	api.GET("/coaching_subjects",handler.GetAll)
	api.PUT("/coaching_subjects/:id", handler.Update)
}