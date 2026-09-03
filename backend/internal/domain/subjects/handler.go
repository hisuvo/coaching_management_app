package subjects

import (
	"coaching_backend/internal/domain/subjects/dto"

	"github.com/labstack/echo/v5"
)

type header struct {
	service Service
}

func NewHeadler(service Service) *header {
	return &header{
		service: service,
	}
}

// POST /api/v1/subjects
// Access: Authencated users (amdin, teacher, manager)
func (h *header) CreateSubject(c *echo.Context) error {
	var req *dto.CreateSubjectRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(3000,map[string]any{
			"success":false,
			"Message": "Invalid request body",
			"Errors":  map[string]string{"error": err.Error()},
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(3000,map[string]any{
			"success":false,
			"Message": "Validation failed",
			"Errors":  map[string]string{"error": err.Error()},
		})
	}

	res, err := h.service.CreateSubject(req)

	if err != nil {
		return c.JSON(500, map[string]any{
			"success":false,
			"Message": "Internal server error",
			"Errors":  map[string]string{"error": err.Error()},
		})
	}
	
	return c.JSON(201, map[string]any{
		"Success": true,
		"Message": "Reservation confirmed successfully",
		"Data":    res,
	})
}