package server

import (
	"coaching_backend/internal/config"
	"coaching_backend/internal/domain/students"
	"coaching_backend/internal/domain/subjects"
	"coaching_backend/internal/domain/users"
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type CustomValidator struct {
  validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
  if err := cv.validator.Struct(i); err != nil {
    // Optionally return the error to let each route control the status code.
    return echo.ErrBadRequest.Wrap(err)
  }
  return nil
}


func Start( db *gorm.DB, cnfg *config.Config) {
	e := echo.New()
	
	e.Validator = &CustomValidator{
		validator: validator.New(),
	}
	
	// e.HTTPErrorHandler = func(c *echo.Context, err error) {
	// 	var appErr *apperror.AppError
	// 	if errors.As(err, &appErr) {
	// 		_ = httpresponse.Error(c, appErr.Status, appErr.Message, appErr.Code)
	// 		return
	// 	}
	// 	echo.DefaultHTTPErrorHandler(false)(c, err)
	// }

	// Middleware
	// e.Use(middleware.RequestLogger())
	// e.Use(middleware.Recover())

	e.GET("/", func(c *echo.Context) error {
        return c.JSON(200, map[string]any{
			"success":true,
			"message": "Hello, World!,Now start go new project.",
			"details":"This is coaching center management project",

		})
    })

	// all route
	users.RegisterRoute(e, db)
	subjects.RegisterRoute(e, db)
	students.RegisterRoute(e, db)

	sc := echo.StartConfig{Address: ":" + cnfg.PORT}
	if err := sc.Start(context.Background(), e); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}