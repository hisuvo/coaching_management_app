package auth

import (
	"coaching_backend/internal/apperror"
	"coaching_backend/internal/httpresponse"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v5"
)

const (
	ContextUserID     = "auth.user_id"
	ContextCoachingID = "auth.coaching_id"
	ContextRole       = "auth.role"
	ContextSessionID  = "auth.session_id"
	ContextUser       = "auth"
)

// Task : authentication
func (h *handler) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {
			return httpresponse.Error(c, http.StatusUnauthorized, "authorization header is required", nil)
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			return httpresponse.Error(c, http.StatusUnauthorized, "invalid authorization header", nil)
		}

		claims, err := h.tokenManger.VerifyAccessToken(parts[1])
		if err != nil {
			return httpresponse.Error(c, http.StatusUnauthorized, "invalid and expired access token", nil)
		}

		if h.authRepository != nil {
			session, err := h.authRepository.GetSessionByID(c.Request().Context(), claims.SessionID)
			if err != nil || session == nil {
				return httpresponse.Error(c, http.StatusUnauthorized, "session not found or invalid", nil)
			}

			if session.RevokedAt != nil || time.Now().After(session.ExpiresAt) {
				return httpresponse.Error(c, http.StatusUnauthorized, "session is revoked or expired", nil)
			}
		}

		c.Set(ContextUserID, claims.UserID)
		c.Set(ContextCoachingID, claims.CoachingID)
		c.Set(ContextRole, claims.Role)
		c.Set(ContextSessionID, claims.SessionID)
		c.Set(ContextUser, claims)

		return next(c)
	}
}

func GetUserRole(c *echo.Context) (string, bool) {
	value := c.Get(ContextRole)
	role, ok := value.(string)

	return role, ok
}

func RequireRoles(roles ...string) echo.MiddlewareFunc {

	allowed := make(map[string]struct{})

	for _, role := range roles {
		allowed[role] = struct{}{}
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c *echo.Context) error {

			role, ok := GetUserRole(c)

			if !ok {
				return apperror.Unauthorized(
					"authentication required",
				)
			}

			if _, exists := allowed[role]; !exists {
				return apperror.Forbidden(
					"you do not have permission",
				)
			}

			return next(c)
		}
	}
}