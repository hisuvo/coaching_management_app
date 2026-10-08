package teachers

import (
	"coaching_backend/internal/domain/teachers/dto"
	"coaching_backend/internal/httpresponse"
	"coaching_backend/internal/pkg/helper"
	"errors"
	"net/http"

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

	coachingID, err := helper.GetCoachingId(c)

	if err != nil {
		return nil
	}

	var req dto.CreateTeacherRequest

	// Bind request body to DTO.
	if err := c.Bind(&req); err != nil {
		return httpresponse.Error(c,http.StatusBadRequest,"invalid request body",err.Error())
	}

	// Validate request body.
	if err := c.Validate(&req); err != nil {
		return httpresponse.Error(c,http.StatusBadRequest,"validation failed",err.Error())
	}

	// Create teacher.
	result, err := h.service.Create(c.Request().Context(), *coachingID, &req)

	if err != nil {
		if errors.Is(err, ErrEmployeeNoExists) {
			return httpresponse.Error(c,http.StatusConflict,"employee number already exists",err.Error())
		}

		return httpresponse.Error(c,http.StatusInternalServerError,"failed to create teacher",err.Error())
	}

	return httpresponse.OK(c,"Teacher created successfully",result)
}

func (h *Handler) GetAll(c *echo.Context) error {
	coachingId, err := helper.GetCoachingId(c)

	if err != nil {
		return err
	}

	result, err := h.service.GetAll(c.Request().Context(), *coachingId)

	if err != nil {
		return httpresponse.Error(c,http.StatusInternalServerError,"teachers retrived failed",err.Error())
	}

	return httpresponse.OK(c, "teachers retrived successfully", result)
}