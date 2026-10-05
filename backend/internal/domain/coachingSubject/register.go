package coachingsubject

import (
	"coaching_backend/internal/domain/auth"
	"coaching_backend/internal/domain/coaching"
	"coaching_backend/internal/domain/subjects"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoute(api *echo.Group, db *gorm.DB, authMiddleware echo.MiddlewareFunc) {
	repository := NewRepository(db)
	subjectRepo := subjects.NewRepository(db)
	coachingRepo := coaching.NewRepository(db)
	service := NewService(repository, subjectRepo, coachingRepo)
	handler := NewHandler(service)

	api.POST("/coaching_subjects", handler.Create, authMiddleware, auth.RequireRoles("SUPER_ADMIN"))
	api.GET("/coaching_subjects",handler.GetOwnCoachingSubjects, authMiddleware)
	api.GET("/coaching_subjects",handler.GetAll)
	api.GET("/coaching_subjects/:id",handler.GetSingleById)
	api.PUT("/coaching_subjects/:id", handler.Update)
	api.DELETE("/coaching_subjects/:id",handler.Delete)
}