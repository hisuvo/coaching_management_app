package subjects

import (
	"coaching_backend/internal/apperror"
	"coaching_backend/internal/domain/subjects/dto"
	"errors"
	"fmt"
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

// GET /api/v1/subjects
// Access: All user can access this
func (h *handler) GetAll(c *echo.Context) error {
	res, err := h.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest,apperror.NotFound("subjects not found"))
	}

	return c.JSON(http.StatusOK, res)
}

// GET /api/v1/:subjectId
// Accesss: only authenticate users
func (h *handler) FindById(c *echo.Context) error {
	id := c.Param("subjectId")
	
	res, err := h.service.FindById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest,apperror.NotFound("subjects Id is not found"))
	}

	return c.JSON(http.StatusOK, res)
}

// PUT /api/v1/:subjectId
// Access: only authenticate user can update her subject
func (h *handler) UpdateSubject(c *echo.Context) error {
	id := c.Param("subjectId")

	var req dto.UpdateSubjectRequest

	if err := c.Bind(&req); err != nil {
		fmt.Println("Bind error:", err)
		return c.JSON(http.StatusBadRequest, apperror.BadRequest("Invalide subject update request"))
	}
	
	if err := c.Validate(&req); err != nil {
		fmt.Println("Validate error:",err)
		return c.JSON(http.StatusBadRequest, apperror.BadRequest("subject update validation failed"))

	}
	
	res, err := h.service.Update(id, &req)

	if err != nil {
		return err
	}

	return  c.JSON(http.StatusCreated, res)
}

// DELETE /api/v1/:subjectId
// Access: only authenticate user can delete this own create subjects
// admin can delete all users subject
func (h *handler) DeleteSubject(c *echo.Context) error {
	id := c.Param("subjectId")
	
	response, err := h.service.Delete(id)

	if err != nil {
		if errors.Is(err, ErrSubjectNotFound) {
			return c.JSON(
				http.StatusNotFound,
				apperror.NotFound("subject not found"),
			)
		}

		return err
	}

	return c.JSON(http.StatusOK, response)
}

// Query /api/v1/subjects?limit=10&page=1&class=8
// only test perpose use this
// todo: user for filter, search and sort subject
func (h *handler) CheckQuery(c *echo.Context) error {
	var query dto.SubjectQuery

	if err := c.Bind(&query); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid query parameters",
		})
	}

	if query.Class == 0 {
		query.Class = 20
	}

	if query.Limit == 0 {
		query.Limit = 10
	}

	if query.Page == 0 {
		query.Page = 1
	}

	return c.JSON(http.StatusOK, query)
}