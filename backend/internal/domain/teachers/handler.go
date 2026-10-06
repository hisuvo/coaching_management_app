package teachers

import (
	"coaching_backend/internal/domain/auth"
	"coaching_backend/internal/domain/teachers/dto"
	"coaching_backend/internal/httpresponse"
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
	// Get coaching ID from authenticated user's context.
	coachingIDValue := c.Get(auth.ContextCoachingID)

	// Make sure coaching ID exists in context.
	if coachingIDValue == nil {
		return httpresponse.Error(
			c,
			http.StatusUnauthorized,
			"coaching information not found",
			nil,
		)
	}

	// Convert context value to uint.
	coachingID, ok := coachingIDValue.(uint)
	if !ok {
		return httpresponse.Error(
			c,
			http.StatusUnauthorized,
			"invalid coaching id",
			nil,
		)
	}

	var req dto.CreateTeacherRequest

	// Bind request body to DTO.
	if err := c.Bind(&req); err != nil {
		return httpresponse.Error(
			c,
			http.StatusBadRequest,
			"invalid request body",
			err.Error(),
		)
	}

	// Validate request body.
	if err := c.Validate(&req); err != nil {
		return httpresponse.Error(
			c,
			http.StatusBadRequest,
			"validation failed",
			err.Error(),
		)
	}

	// Create teacher.
	result, err := h.service.Create(
		c.Request().Context(),
		coachingID,
		&req,
	)

	if err != nil {
		if errors.Is(err, ErrEmployeeNoExists) {
			return httpresponse.Error(
				c,
				http.StatusConflict,
				"employee number already exists",
				err.Error(),
			)
		}

		return httpresponse.Error(
			c,
			http.StatusInternalServerError,
			"failed to create teacher",
			err.Error(),
		)
	}

	return httpresponse.OK(
		c,
		"Teacher created successfully",
		result,
	)
}