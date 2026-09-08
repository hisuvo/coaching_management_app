package coaching

import (
	"coaching_backend/internal/domain/coaching/dto"
	"coaching_backend/internal/httpresponse"
	"net/http"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler{
	return &handler{
		service: service,
	}
}

func (h *handler) Create(c *echo.Context) error {
	var req dto.CreateCoachingRequest

	if err := c.Bind(&req); err != nil {
		return httpresponse.Error(c, http.StatusBadRequest,"Invalide caching create request",err.Error())
	}

	if err := c.Validate(&req); err != nil {
		return httpresponse.Error(c, http.StatusBadRequest,"coaching create validation failed",err.Error())
	}

	coaching, err := h.service.Create(&req)

	if err != nil {
		return httpresponse.Error(c, http.StatusBadRequest,"coaching generate failed",err.Error())
	}

	return httpresponse.OK(c, "coaching generate successful", coaching)
}

// GET api/v1/coaching/:id
// Access: only authenticate user can access
func (h *handler) GetById(c *echo.Context) error {
	id := c.Param("id")

	result, err := h.service.GetById(id)

	if err != nil {
		return httpresponse.Error(c,http.StatusBadGateway, "Coaching info retrive fail", err.Error())
	}

	return httpresponse.OK(c, "Coaching info retived successfuly", result)
}

// GET api/v1/coaching
func (h *handler) GetAll(c *echo.Context) error {
	result, err := h.service.GetAll()

	if err != nil {
		return httpresponse.Error(c, http.StatusNotFound ,"Coaching data retirved failed", err.Error())
	}

	return httpresponse.OK(c, "All coaching data retrived successfully", result)
}