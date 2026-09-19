package assignments

import (
	"coaching_backend/internal/domain/assignments/dto"
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
	var req dto.CreateAssignmentRequest

	if err := c.Bind(&req); err != nil {
		return httpresponse.Error(c,http.StatusBadRequest,"Invalide assignment request", err.Error())
	}

	if err := c.Validate(&req); err != nil {
		return httpresponse.Error(c,http.StatusBadRequest,"assignment validation failed", err.Error())
	}

	// Example:
	// Get user ID from JWT middleware/context.
	createdBy := uint(1)

	assignment, err := h.service.Create(&req, createdBy)

	if err != nil {
		return httpresponse.Error(c,http.StatusBadRequest,"Failde to create assignment", err.Error())
	}

	return httpresponse.OK(c,"Successfully assignment created done!", assignment)
}

func (h *Handler) GetAll(c *echo.Context) error {
	// var assignments []*dto.AssignmentResponse
	assignments, err := h.service.GetAll()

	if err != nil {
		return httpresponse.Error(c, http.StatusBadRequest, "Assignment data retrived failed", err.Error())
	}

	return httpresponse.OK(c,"Assignments data retrived successfuly", assignments)
}

func (h *Handler) GetById(c *echo.Context) error {
	assignmentId := c.Param("assignmentId")

	assignment, err := h.service.GetById(assignmentId)

	if err != nil {
		return httpresponse.Error(c, http.StatusBadRequest, "your assignment retrived failed", err.Error())
	}

	return httpresponse.OK(c,  "your assignment retrived successfull", assignment)

}