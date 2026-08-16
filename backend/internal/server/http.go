package server

import (
	"coaching_backend/internal/config"

	"github.com/labstack/echo/v5"
)

func Start() {
	e := echo.New()
	cfg := config.EnvConfig()

	e.GET("/", func(c *echo.Context) error {
        return c.JSON(200, map[string]string{
			"message": "Hello, World!,Now start go new project.",
			"Port":cfg.Port,
			"details":"This is coaching center management project",
		})
    })

	e.Start(":5000")
}