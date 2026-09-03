package subjects

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoute(e *echo.Echo, db *gorm.DB) {
	userRepo := NewRepository(db)
	userService := NewService(userRepo)
	userHandler := NewHeadler(userService)

	subjects := e.Group("/api/v1/subjects")

	subjects.POST("/", userHandler.CreateSubject)
}