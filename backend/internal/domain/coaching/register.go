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

	route := e.Group("/api/v1", authMiddleware)

	admin := route.Group("")
	admin.Use((auth.RequireRoles("ADMIN")))
	admin.GET("/coachings/:id",coachingHandler.GetById)

	memeber := route.Group("/m",auth.RequireRoles("MEMBER"))
	memeber.POST("/coachings", coachingHandler.Create)
	memeber.GET("/coachings", coachingHandler.GetAll)
}