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
	
	protectedRoute := route.Group("", authMiddleware)
	protectedRoute.POST("/coachings", coachingHandler.Create, auth.RequireRoles("SUPER_ADMIN"))
	protectedRoute.DELETE("/coachings/:id",coachingHandler.Delete, auth.RequireRoles("SUPER_ADMIN"))
	protectedRoute.PUT("/coachings/:id",coachingHandler.Update, auth.RequireRoles("ADMIN"))

	// ------ NOTE ------
	// protected := route.Group("")
	// protected.Use(authMiddleware)
	// protected.Use((auth.RequireRoles("ADMIN","SUPER_ADMIN")))
	// protected.POST("/coachings", coachingHandler.Create)
	// protected.DELETE("/coachings/:id",coachingHandler.Delete)
	// protected.PUT("/coachings/:id",coachingHandler.Update)
}