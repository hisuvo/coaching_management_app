package coachingsubject

import (
	"coaching_backend/internal/httpresponse"

	"github.com/labstack/echo/v5"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Create(c *echo.Context) error {
	return httpresponse.OK(c, "coaching_subject create successfully", nil)
}

func (h *Handler) GetAll(c *echo.Context) error {
	return httpresponse.OK(c,"coaching_subject data retrived successfully", nil)
}

func (h *Handler) Update(c *echo.Context) error {
	return httpresponse.OK(c,"coaching_subject data updated successfully", nil)
}