package coaching

import (
	"coaching_backend/internal/domain/coaching/dto"
	"coaching_backend/internal/httpresponse"
	"errors"
	"net/http"
	"strconv"

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

// POST : api/v1/coachings
// Access: only super-admin can generate coaching
func (h *handler) Create(c *echo.Context) error {
	var req dto.CreateCoachingRequest

	if err := c.Bind(&req); err != nil {
		return httpresponse.Error(c, http.StatusBadRequest,"Invalide caching create request",err.Error())
	}

	if err := c.Validate(&req); err != nil {
		return httpresponse.Error(c, http.StatusBadRequest,"coaching create validation failed",err.Error())
	}

	coaching, err := h.service.Create(c.Request().Context(),&req)

	if err != nil {
		return httpresponse.Error(c, http.StatusBadRequest,"coaching generate failed",err.Error())
	}

	return httpresponse.OK(c, "coaching generate successful", coaching)
}

// GET api/v1/coachings/:id
// Access: only authenticate user can access
func (h *handler) GetById(c *echo.Context) error {
	coachingId := c.Param("id")

	id, err := strconv.ParseInt(coachingId, 10, 64)
	if err != nil {
		return httpresponse.Error(c,http.StatusBadRequest,"Invalided coaching id",nil)
	}

	result, err := h.service.GetById(c.Request().Context(), uint(id))

	if err != nil {
		return httpresponse.Error(c,http.StatusBadGateway, "Coaching info retrive fail", err.Error())
	}

	return httpresponse.OK(c, "Coaching info retived successfuly", result)
}

// GET api/v1/coachings
// Access: All user can show this coaching
func (h *handler) GetAll(c *echo.Context) error {
	result, err := h.service.GetAll(c.Request().Context())

	if err != nil {
		return httpresponse.Error(c, http.StatusNotFound ,"Coaching data retirved failed", err.Error())
	}

	return httpresponse.OK(c, "All coaching data retrived successfully", result)
}

// PUT api/v1/coachings
// Access: Only SuperAdmin can
func (h *handler) Update(c *echo.Context) error {
	ctx := (*c).Request().Context()

	coachingId := (*c).Param("id")
	
	id, err := strconv.ParseInt(coachingId, 10, 64)
	
	if  err != nil {
		return httpresponse.Error(c,http.StatusBadRequest,"invalid coaching id", err.Error())
	}

	var req dto.UpdateCoachingRequest

	if err := (*c).Bind(&req); err != nil {
		return httpresponse.Error(c,http.StatusBadRequest,"invalid request body", err.Error())
	}

	if err := (*c).Validate(&req); err != nil {
		return httpresponse.Error(c,http.StatusBadRequest,"coaching update validation failed", err.Error())
	}
	
	response, err := h.service.Update(ctx, uint(id), &req)

	if err != nil {
		if errors.Is(err,ErrCoachingNotFound){
			return httpresponse.Error(c,http.StatusBadRequest,"Coaching not found",err.Error())
		}
		return httpresponse.Error(c,http.StatusBadRequest,"Coaching update failed",err.Error())
	}

	return httpresponse.Error(c,http.StatusOK,"Coaching updated successfully",response)
}

// DELETE api/v1/coachings
// Access: Only Super-Admin can delete
func (h *handler) Delete(c *echo.Context) error {
	coachingId := (*c).Param("id")

	id, err := strconv.ParseInt(coachingId, 10, 64)

	if err != nil {
		return httpresponse.Error(c,http.StatusBadRequest,"Invalied coaching id", err.Error())
	}

	if err := h.service.Delete(c.Request().Context(), uint(id)); err != nil {
		return httpresponse.Error(c,http.StatusBadRequest,"Coaching delete failed", err.Error())
	}

	return httpresponse.Error(c,http.StatusOK,"Coaching delete successfully", nil)
}