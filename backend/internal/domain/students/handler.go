package students

import (
	"coaching_backend/internal/apperror"
	"coaching_backend/internal/domain/students/dto"
	"coaching_backend/internal/httpresponse"
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

	var req dto.CreateStudentRequest
	
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, apperror.BadRequest("Invalide student create request"))
	}


	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, apperror.BadRequest("student createing validation failed"))
	}

	student, err := h.service.Create(&req)

    if err != nil {
        return httpresponse.Error(c, http.StatusInternalServerError,"student creation failed", err.Error())
    }

	return httpresponse.OK(c,"student create successful", student)
}

