package coaching

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoute(e *echo.Echo, db *gorm.DB){
	coachingRepo := NewRepository(db)
	coachingService := NewService(coachingRepo)
	coachingHandler := NewHandler(coachingService)

	route := e.Group("/api/v1")

	route.POST("/coachings", coachingHandler.Create)
	route.GET("/coachings/:id",coachingHandler.GetById)
	route.GET("/coachings", coachingHandler.GetAll)
}