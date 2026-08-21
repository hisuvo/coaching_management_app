package httpresponse

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func Success(c *echo.Context, status int, message string, data interface{}) error {
	return c.JSON(status, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Error(c *echo.Context, status int, message string, errors interface{}) error {
	return c.JSON(status, Response{
		Success: false,
		Message: message,
		Errors:  errors,
	})
}

func OK(c *echo.Context, message string, data interface{}) error {
	return Success(c, http.StatusOK, message, data)
}

func Created(c *echo.Context, message string, data interface{}) error {
	return Success(c, http.StatusCreated, message, data)
}

// instead of : return c.JSON(http.StatusOK, user)
// use : return httpresponse.OK(c, "User fetched successfully", user)