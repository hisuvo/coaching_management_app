package server

import (
	"coaching_backend/internal/apperror"
	"coaching_backend/internal/config"
	"coaching_backend/internal/domain/users"
	"coaching_backend/internal/httpresponse"
	"context"
	"errors"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/gorm"
)

func Start( db *gorm.DB, cnfg *config.Config) {
	e := echo.New()
	
	e.HTTPErrorHandler = func(c *echo.Context, err error) {
		var appErr *apperror.AppError
		if errors.As(err, &appErr) {
			_ = httpresponse.Error(c, appErr.Status, appErr.Message, appErr.Code)
			return
		}
		echo.DefaultHTTPErrorHandler(false)(c, err)
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

	// all route
	users.RegisterRoute(e, db)
	// subjects.

	sc := echo.StartConfig{Address: ":" + cnfg.PORT}
	if err := sc.Start(context.Background(), e); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}