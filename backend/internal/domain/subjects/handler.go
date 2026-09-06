package subjects

import (
	"coaching_backend/internal/apperror"
	"coaching_backend/internal/domain/subjects/dto"
	"net/http"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

// POST /api/v1/subjects
// Access: Authencated users (amdin, teacher, manager)
func (h *handler) CreateSubject(c *echo.Context) error {
	var req dto.CreateSubjectRequest

	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	response, err := h.service.CreateSubject(&req)

	if err != nil {
	return err
}
	return c.JSON(http.StatusCreated, response)
}

func (h *handler) GetAll(c *echo.Context) error {
	res, err := h.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest,apperror.NotFound("subjects not found"))
	}

	return c.JSON(http.StatusOK, res)
}