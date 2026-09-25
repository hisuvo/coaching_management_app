package branches

import (
	"coaching_backend/internal/domain/auth"
	"coaching_backend/internal/domain/branches/dto"
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
		return httpresponse.Error(c, http.StatusBadRequest, "Invalid branch request", err.Error())
	}

	branch, err := h.service.Create(c.Request().Context(), &req)
	if err != nil {
		return httpresponse.Error(c, http.StatusInternalServerError, "Failed to create branch", err.Error())
	}

	return httpresponse.OK(c, "Branch created successfully", branch)
}