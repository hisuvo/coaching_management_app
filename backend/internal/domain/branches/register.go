package branches

import (
	"fmt"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoute(e *echo.Echo, db *gorm.DB) {

	route := e.Group("/api/v1")

	route.GET("/branches", func(c *echo.Context) error {
		fmt.Println("Hello world form branch")
		return nil
	} )
}