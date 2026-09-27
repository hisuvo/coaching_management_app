package branches

import (
	"coaching_backend/internal/domain/auth"
	"coaching_backend/internal/domain/branches/dto"
	"coaching_backend/internal/httpresponse"
	"coaching_backend/internal/pkg"
	"net/http"
	"strconv"

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

func (h *Handler) CreateBranch(c *echo.Context) error {

	coachingID, ok := c.Get(auth.ContextCoachingID).(*uint)
	
	if !ok || coachingID == nil {
		return httpresponse.Error(c, http.StatusUnauthorized, "coaching id not found", nil)
	}

	var req dto.CreateBranchRequest

	if *coachingID > 0 {
		req.CoachingID = *coachingID
	}

	if err := c.Bind(&req); err != nil {
		return httpresponse.Error(c, http.StatusBadRequest, "Invalid branch request", err.Error())
	}

	if err := c.Validate(&req); err != nil {
		return httpresponse.Error(c, http.StatusBadRequest, "Invalid branch request", pkg.FormatValidationErrors(err))
	}

	branch, err := h.service.Create(c.Request().Context(), &req)
	if err != nil {
		return httpresponse.Error(c, http.StatusInternalServerError, "Failed to create branch", err.Error())
	}

	return httpresponse.OK(c, "Branch created successfully", branch)
}

func (h *Handler) FindByID(c *echo.Context) error {
	branchId := (*c).Param("id")
	id, err := strconv.ParseInt(branchId, 10, 64)

	if err != nil {
		return httpresponse.Error(c, http.StatusBadGateway, "brnach id is not valied", err.Error())
	}

	branch, err := h.service.FindByID((*c).Request().Context(), uint(id))

	if err != nil {
		return httpresponse.Error(c, http.StatusBadGateway, "Brnach retrived failed", err.Error())
	}

	return httpresponse.OK(c, "Branch retrived successfully", branch)
}

func (h *Handler) FindAll(c *echo.Context) error {

	branchs, err := h.service.FindAll(c.Request().Context())

	if err != nil {
		return httpresponse.Error(c,http.StatusBadRequest,"branches retived failed", err.Error())
	}

	return httpresponse.OK(c,"branches retived successfully", branchs)
}

func (h *Handler) Update(c *echo.Context) error {

	// get branch id form url paramater
	branchIDParam := c.Param("id")

	// convert url parameter form stirng to unit
	branchID, err := strconv.ParseUint(branchIDParam, 10, 64)

	if err != nil {
		return httpresponse.Error(c, http.StatusBadRequest, "invalied branch id", err.Error())
	}


	var req dto.UpdateBranchRequest

	// Bind request body into the dto
	if err := c.Bind(&req); err != nil {
		return httpresponse.Error( c, http.StatusUnsupportedMediaType, "invalid branch update request", err.Error())
	}
	if err := c.Validate(&req); err != nil {
		return httpresponse.Error(c, http.StatusBadGateway, "branch update validation failed", err.Error())
	}

	// call service layer
	branch, err := h.service.Update(c.Request().Context(), uint(branchID), &req)

	if err != nil {
		return httpresponse.Error(c, http.StatusBadRequest, "branch update failed", err.Error())
	}

	return httpresponse.OK(c,"branch updated successfully", branch)
}

func (h *Handler) Delete(c *echo.Context) error {
	// get branch id form url paramater
	branchIDParam := c.Param("id")

	// convert url parameter form stirng to unit
	branchID, err := strconv.ParseUint(branchIDParam, 10, 64)

	if err != nil {
		return httpresponse.Error(c, http.StatusBadRequest, "invalied branch id", err.Error())
	}

	// call service layer
	branch, err := h.service.Delete(c.Request().Context(), uint(branchID))

	if err != nil {
		return httpresponse.Error(c, http.StatusBadRequest, "branch deleted failed", err.Error())
	}

	return httpresponse.OK(c,"branch deleted successfully", branch)
}