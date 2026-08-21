package server

import (
	"coaching_backend/internal/config"
	"context"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/gorm"
)

func Start( db *gorm.DB, cnfg *config.Config) {
	e := echo.New()

	// Middleware
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.GET("/", func(c *echo.Context) error {
        return c.JSON(200, map[string]string{
			"message": "Hello, World!,Now start go new project.",
			"details":"This is coaching center management project",
		})
    })


	sc := echo.StartConfig{Address: ":" + cnfg.PORT}
	if err := sc.Start(context.Background(), e); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}