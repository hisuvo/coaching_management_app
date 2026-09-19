package users

import (
	"coaching_backend/internal/apperror"
	"coaching_backend/internal/domain/users/dto"
	"coaching_backend/internal/httpresponse"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

// type Handler interface{}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Register(c *echo.Context) error {
	var req dto.CreateUserRequest

	if err := c.Bind(&req); err != nil {
		return httpresponse.Error(c, http.StatusBadRequest,"Invalide account create request", map[string]string{
				"code":    "INVALID_REQUEST",
				"message": "Request body is invalid.",
			},)
	}

	user, err := h.service.Register(c.Request().Context(),req)


	if err != nil {
		// Duplicate email
		if errors.Is(err, ErrDuplicateEmail){
			return httpresponse.Error(c, http.StatusConflict,"User createion failed", map[string]string{
				"code":"DUPLLICATED_EMAIL",
				"message": "Email address is already in use.",
			})
		}

		// record not found
		if errors.Is(err, gorm.ErrRecordNotFound){
			return httpresponse.Error(c, http.StatusNotFound,"User createion failed", map[string]string{
				"code": "NOT_FOUND",
				"message": "Required resource was not found.",
			})
		}

		// Unknown/internal error
		return httpresponse.Error(c, http.StatusInternalServerError,"User createion failed", map[string]string{
				"code":    "INTERNAL_SERVER_ERROR",
				"message": "Something went wrong. Please try again later.",
			})
	}

	return httpresponse.OK(c,"Account created successfuly",user)
}

func (h *Handler) FindByEmail(c *echo.Context) error {
	email := c.Param("email")

	if email == "" {
		return apperror.BadRequest("Email is required")
	}

	user, err := h.service.FindByEmail(c.Request().Context(), email)

	if err != nil {
		// record not found
		if errors.Is(err, ErrEmailNotFound){
			return httpresponse.Error(c, http.StatusNotFound,"User createion failed", map[string]string{
				"code": "NOT_FOUND",
				"message": "Required resource was not found.",
			})
		}

		return err
	}

	return httpresponse.Error(c, http.StatusNotFound,"User createion failed", user)
}

func (h *Handler) GetByID(c *echo.Context) error{
	idStr := c.Param("id")

	if idStr == "" {
		return apperror.BadRequest("User id is required")
	}

	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		return httpresponse.Error(c,http.StatusBadRequest,
			"User ID must be a valid number.",
			err.Error(),)
	}

	user, err := h.service.GetByID(c.Request().Context(), uint(id))

	if err != nil {
		if errors.Is(err, ErrUserNotRound){
			return httpresponse.Error(c,http.StatusBadRequest,
			"User not found!",
			err.Error(),)
		}
	}

	return httpresponse.OK(c,"User retrived successfuly",user)
}