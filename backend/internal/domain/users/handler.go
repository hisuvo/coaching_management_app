package users

import (
	"coaching_backend/internal/apperror"
	"coaching_backend/internal/domain/users/dto"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
)

// type Handler interface{}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Register(c *echo.Context) error {
	var req dto.CreateUserRequest

	if err := c.Bind(&req); err != nil {
		return err
	}

	user, err := h.service.Register(req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *Handler) FindByEmail(c *echo.Context) error {
	email := c.Param("email")

	if email == "" {
		return apperror.BadRequest("Email is required")
	}

	user, err := h.service.FindByEmail(email)

	if err != nil {
		return err
	}

	fmt.Println(user)

	return c.JSON(http.StatusOK, user)
}