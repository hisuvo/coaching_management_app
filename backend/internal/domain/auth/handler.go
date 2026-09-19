package auth

import (
	"coaching_backend/internal/config"
	"coaching_backend/internal/domain/auth/dto"
	"coaching_backend/internal/httpresponse"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
)

// Task : HTTP endpoints

type handler struct {
	service Service
	config config.Config
}

func NewHandler(service Service, config config.Config) *handler{
	return &handler{
		service: service,
		config: config,
	}
}

func (h *handler) setRefreshCookie(c *echo.Context, token string, expiresAt time.Time,) {

	cookie := new(http.Cookie)

	cookie.Name = h.config.COOKIE_NAME
	cookie.Value = token
	cookie.Path = "/"
	cookie.HttpOnly = false
	cookie.Secure = h.config.COOKIE_SECURE
	cookie.SameSite = http.SameSiteLaxMode
	cookie.Expires = expiresAt

	cookie.MaxAge = int(
		time.Until(expiresAt).Seconds(),
	)


	http.SetCookie(c.Response(), cookie)
}

func (h *handler) clearRefreshCookie(c *echo.Context,) {

	cookie := new(http.Cookie)

	cookie.Name = h.config.COOKIE_NAME
	cookie.Value = ""
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Secure = h.config.COOKIE_SECURE
	cookie.MaxAge = -1
	cookie.Expires = time.Unix(0, 0)

	http.SetCookie(c.Response(), cookie)
}

func (h *handler) Login(c *echo.Context) error {
	var req dto.LoginRequest

	if err := c.Bind(&req); err != nil {
		return httpresponse.Error(c, http.StatusBadRequest,"Invalid request body", err.Error()) 
	}

	if err := c.Validate(&req); err != nil {
		return httpresponse.Error(c, http.StatusBadRequest,"validation failed", err.Error()) 
	}

	result, err := h.service.Login(c.Request().Context(), &req, c.Request().UserAgent(), c.RealIP())

	if err != nil {
		return httpresponse.Error(c, http.StatusUnauthorized,"invalid email or password", err.Error()) 
	}

	expiresAt := time.Now().Add(
		h.config.JWT_REFRESH_EXPIRES_IN,
	)

	h.setRefreshCookie(
		c,
		result.RefreshToken,
		expiresAt,
	)

	response := dto.LoginResponse{
		AccessToken: result.AccessToken,
		TokenType:   "Bearer",
		ExpiresIn:   result.ExpiresIn,

		User: dto.UserResponse{
			ID:         result.UserID,
			CoachingID: result.CoachingID,
			Name:       result.Name,
			Email:      result.Email,
			Role:       result.Role,
			Status:     "ACTIVE",
		},
	}


	return httpresponse.OK(c,"Login successful",response)
}