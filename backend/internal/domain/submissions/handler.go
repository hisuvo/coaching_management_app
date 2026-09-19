package submissions

import (
	"coaching_backend/internal/domain/submissions/dto"
	"coaching_backend/internal/httpresponse"
	"net/http"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service Service
}

func NewHandler(service Service) handler {
	return handler{
		service: service,
	}
}

func (h *handler) Create(c *echo.Context) error {
	var submission dto.CreateSubmissionRequest

	if err := c.Bind(&submission); err != nil {
		return httpresponse.Error(c, http.StatusBadRequest,"Invalide request", err.Error())
	}
	
	if err := c.Validate(&submission); err != nil {
		return httpresponse.Error(c, http.StatusBadRequest,"Validation error", err.Error())
	}


	res, err := h.service.Create("12", &submission)

	if err != nil {
		return httpresponse.Error(c,http.StatusBadRequest,"submission creating failed", err.Error())
	}

	return c.JSON(201, res)
}