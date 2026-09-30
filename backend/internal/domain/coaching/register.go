package coaching

import (
	"coaching_backend/internal/domain/auth"
	"fmt"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoute(api *echo.Group, db *gorm.DB, authMiddleware echo.MiddlewareFunc){
	coachingRepo := NewRepository(db)
	coachingService := NewService(coachingRepo)
	coachingHandler := NewHandler(coachingService)


	api.GET("/coachings", coachingHandler.GetAll)
	api.GET("/coachings/:id",coachingHandler.GetById)
	api.GET("/coaching/:id/admin",func(c *echo.Context) error {
		fmt.Println("coaching admin info")
		return nil
	})
	api.GET("/coaching/:id/details",func(c *echo.Context) error {
		fmt.Println("coaching details")
		return nil
	})
	api.POST("/coachings", coachingHandler.Create, authMiddleware, auth.RequireRoles("PLATFORM_ADMIN"))
	api.DELETE("/coachings/:id",coachingHandler.Delete, authMiddleware, auth.RequireRoles("PLATFORM_ADMIN"))
	api.PUT("/coachings/:id",coachingHandler.Update, authMiddleware, auth.RequireRoles("PLATFORM_ADMIN"))

	// ------ NOTE ------
	// protected := route.Group("")
	// protected.Use(authMiddleware)
	// protected.Use((auth.RequireRoles("ADMIN","SUPER_ADMIN")))
	// protected.POST("/coachings", coachingHandler.Create)
	// protected.DELETE("/coachings/:id",coachingHandler.Delete)
	// protected.PUT("/coachings/:id",coachingHandler.Update)
}