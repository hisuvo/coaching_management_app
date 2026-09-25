package branches

import (
	"fmt"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoute(api *echo.Group, db *gorm.DB, authMiddleware echo.MiddlewareFunc) {

	branchRepo := NewRepository(db)
	branchService := NewService(branchRepo)
	branchHandler := NewHandler(branchService)

	api.POST("/branches", branchHandler.CreateBranch, authMiddleware)

	api.GET("/branches", func(c *echo.Context) error {
		fmt.Println("Hello world form branch")
		return c.JSON(200, map[string]any{
			"success": true,
			"message": "Branch fetched successfully",
			"details": "This is branch API",
		})
	})

	api.PUT("/branches/:id", func(c *echo.Context) error {
		fmt.Println("Hello world form branch")
		return c.JSON(200, map[string]any{
			"success": true,
			"message": "Branch updated successfully",
			"details": "This is branch API",
		})
	})

	api.DELETE("/branches/:id", func(c *echo.Context) error {
		fmt.Println("Hello world form branch")
		return c.JSON(200, map[string]any{
			"success": true,
			"message": "Branch deleted successfully",
			"details": "This is branch API",
		})
	})
}