package branches

import (
	"coaching_backend/internal/domain/auth"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoute(api *echo.Group, db *gorm.DB, authMiddleware echo.MiddlewareFunc) {

	branchRepo := NewRepository(db)
	branchService := NewService(branchRepo)
	branchHandler := NewHandler(branchService)

	api.POST("/branches", branchHandler.CreateBranch, authMiddleware, auth.RequireRoles("SUPER_ADMIN","ADMIN"))
	api.GET("/branches", branchHandler.FindAll)
	api.GET("/branches/:id", branchHandler.FindByID)
	api.GET("/coaching_by_branches", branchHandler.FindAll)
	api.PUT("/branches/:id", branchHandler.Update,authMiddleware, auth.RequireRoles("SUPER_ADMIN","ADMIN"))
	api.DELETE("/branches/:id", branchHandler.Delete)
}