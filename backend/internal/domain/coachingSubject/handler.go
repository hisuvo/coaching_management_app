package coachingsubject

import (
	"coaching_backend/internal/domain/auth"
	"coaching_backend/internal/domain/coachingSubject/dto"
	"coaching_backend/internal/httpresponse"
	"net/http"
	"strconv"
	"strings"

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

// if exists then create only coaching_subject fild
// other wise create both subject and and coaching_subject field
func (h *Handler) Create(c *echo.Context) error {

	var req dto.CreateCoachingSubjectRequest

	coachingId := c.Get(auth.ContextCoachingID).(*uint)

	req.CoachingID = *coachingId

	if err := c.Bind(&req); err != nil {
		return httpresponse.Error(c, http.StatusBadRequest, "invalided coaching subject create request", err.Error())
	}

	if err := c.Validate(&req); err != nil {
		return httpresponse.Error(c, http.StatusBadRequest, "validation failed coaching_subject create request", err.Error())
	}

	result, err := h.service.Create(c.Request().Context(), &req)

	if err != nil {
		return httpresponse.Error(c, http.StatusBadRequest, "coaching subject created failed", err.Error())
	}

	return httpresponse.OK(c, "coaching_subject create successfully", result)
}

func (h *Handler) GetAll(c *echo.Context) error {
	result, err := h.service.GetAll(c.Request().Context())

	if err != nil {
		return httpresponse.Error(c,http.StatusBadRequest, "coaching_subjects retived failed", err.Error())
	}

	return httpresponse.OK(c,"coaching_subject data retrived successfully", result)
}

func (h *Handler) GetOwnCoachingSubjects(c *echo.Context) error {
	coaching_id := c.Get(auth.ContextCoachingID).(*uint)

	result, err := h.service.GetOwnCoachingSubject(c.Request().Context(), coaching_id)

	if err != nil {
		return httpresponse.Error(c, http.StatusBadRequest, "coaching subject created failed", err.Error())
	}

	return httpresponse.OK(c, "coaching_subject create successfully", result)
}

func (h *Handler) Update(c *echo.Context) error {
	// Get ID from URL
	coaching_subject_id := c.Param("id")

	id, err := strconv.ParseUint(coaching_subject_id, 10 , 64)

	if err != nil {
		return httpresponse.Error(c, http.StatusBadRequest,"invalid coaching_subject id", err.Error())
	}

	// Parse request body.
	var req dto.UpdateCoachingSubjectRequest

	if err := c.Bind(&req); err != nil {
		return httpresponse.Error(c, http.StatusBadRequest,"invalid coaching subject request", err.Error())
	}

	// validation request
	if err := c.Validate(&req); err != nil {
		return httpresponse.Error(c, http.StatusBadRequest,"validation filed coaching subject request", err.Error())
	}

	status := strings.ToUpper(*req.Status)

	response, err := h.service.Update(c.Request().Context(), uint(id), &status)

	if err != nil {
		return httpresponse.Error(c, http.StatusBadRequest,"coaching_subject updated failed", err.Error())
	}

	// Return successful response.
	return httpresponse.Success(c,http.StatusOK,"coaching subject updated successfully",response)
}

func (h *Handler) GetSingleById(c *echo.Context) error {
	// Get ID from URL
	coaching_subject_id := c.Param("id")

	id, err := strconv.ParseUint(coaching_subject_id, 10 , 64)

	if err != nil {
		return httpresponse.Error(c, http.StatusBadRequest,"invalid coaching_subject id", err.Error())
	}

	coaching_subject, err := h.service.GetSingleById(c.Request().Context(),uint(id))

	return httpresponse.OK(c, "coaching_subject retrived successfully", coaching_subject)
}

func (h *Handler) Delete(c *echo.Context) error {
	// Get ID from URL
	coaching_subject_id := c.Param("id")

	id, err := strconv.ParseUint(coaching_subject_id, 10 , 64)

	if err != nil {
		return httpresponse.Error(c, http.StatusBadRequest,"invalid coaching_subject id", err.Error())
	}

	if err := h.service.Delete(c.Request().Context(), uint(id)); err != nil {
		return err
	}
	return httpresponse.OK(c,"coaching_subject deleted successfully",nil)
}