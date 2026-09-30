package branches

import (
	"coaching_backend/internal/domain/auth"
	"fmt"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoute(api *echo.Group, db *gorm.DB, authMiddleware echo.MiddlewareFunc) {

	branchRepo := NewRepository(db)
	branchService := NewService(branchRepo)
	branchHandler := NewHandler(branchService)

	api.POST("/branches", branchHandler.CreateBranch, authMiddleware, auth.RequireRoles("SUPER_ADMIN","BRANCH_ADMIN"))
	api.GET("/branches", branchHandler.FindAll)
	api.GET("/branches/:id", branchHandler.FindByID)
	api.GET("/branch/:id/amdin",func(c *echo.Context) error {
		fmt.Println("Branch amdin")
		return nil
	})
	api.GET("/branch/:id/details",func(c *echo.Context) error {
		fmt.Println("Branch details")
		return nil
	})
	api.GET("/coaching_by_branches", branchHandler.FindAll)
	api.PUT("/branches/:id", branchHandler.Update, authMiddleware, auth.RequireRoles("SUPER_ADMIN","BRANCH_ADMIN"))
	api.DELETE("/branches/:id", branchHandler.Delete, authMiddleware, auth.RequireRoles("SUPER_ADMIN","BRANCH_ADMIN"))
}