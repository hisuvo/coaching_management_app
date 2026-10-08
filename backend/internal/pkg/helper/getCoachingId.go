package helper

import (
	"coaching_backend/internal/domain/auth"
	"coaching_backend/internal/httpresponse"
	"net/http"

	"github.com/labstack/echo/v5"
)

func GetCoachingId(c *echo.Context) (coachingID *uint, err error) {
	// Get coaching id from authenticated users context
	coachingValue := c.Get(auth.ContextCoachingID)

	// Make sure coaching ID exists in context
	if coachingValue == nil {
		return nil, httpresponse.Error(c, http.StatusUnauthorized, "coaching information not found", nil)
	}

	// Convert context value to uint
	coachingID, ok := coachingValue.(*uint)

	if !ok {
		return nil, httpresponse.Error(c,http.StatusUnauthorized,"invalid coaching id",nil)
	}

	return coachingID, nil
}

