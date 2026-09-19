package users

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoute(e *echo.Echo, db *gorm.DB) {

	userRepo := NewRepository(db)
	userService := NewService(userRepo)
	userHandler := NewHandler(userService)

	users := e.Group("/api/v1")

	users.POST("/auth/register",userHandler.Register)

	users.GET("/users/:email", userHandler.FindByEmail)

	users.GET("/users/:id", userHandler.GetByID)
}