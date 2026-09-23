package coaching

import (
	"coaching_backend/internal/domain/auth"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoute(e *echo.Echo, db *gorm.DB, authMiddleware echo.MiddlewareFunc){
	coachingRepo := NewRepository(db)
	coachingService := NewService(coachingRepo)
	coachingHandler := NewHandler(coachingService)

	route := e.Group("/api/v1")
	route.GET("/coachings", coachingHandler.GetAll)
	route.GET("/coachings/:id",coachingHandler.GetById)
	route.PUT("/coachings/:id",coachingHandler.Update)

	admin := route.Group("")
	admin.Use(authMiddleware)
	admin.Use((auth.RequireRoles("ADMIN","SUPER_ADMIN")))
	admin.POST("/coachings", coachingHandler.Create)
	
}