package server

import (
	"coaching_backend/internal/config"
	"coaching_backend/internal/routes"
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/gorm"
)

type CustomValidator struct {
  validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
  if err := cv.validator.Struct(i); err != nil {
    return echo.ErrBadRequest.Wrap(err)
  }
  return nil
}


func Start( db *gorm.DB, cnfg *config.Config) {
	e := echo.New()
	
	e.Validator = &CustomValidator{
		validator: validator.New(),
	}

	// Middleware
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.GET("/", func(c *echo.Context) error {
        return c.JSON(200, map[string]any{
			"success":true,
			"message": "Hello, World!,Now start go new project.",
			"details":"This is coaching center management project",

		})
    })

	routes.RegisterRoutes(e, db, cnfg)

	sc := echo.StartConfig{Address: ":" + cnfg.PORT}
	if err := sc.Start(context.Background(), e); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}