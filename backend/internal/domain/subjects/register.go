package subjects

import (
	"coaching_backend/internal/domain/auth"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoute(api *echo.Group, db *gorm.DB, authMiddleware echo.MiddlewareFunc) {
	subjectRepo := NewRepository(db)
	subjectService := NewService(subjectRepo)
	subjectHandler := NewHandler(subjectService)

	api.POST("/subjects", subjectHandler.CreateSubject, authMiddleware, auth.RequireRoles("SUPER_ADMIN"))
	api.GET("/subjects", subjectHandler.GetAll)
	api.GET("/subjects/query", subjectHandler.CheckQuery)
	api.GET("/subjects/:subjectId", subjectHandler.GetById, authMiddleware, auth.RequireRoles("SUPER_ADMIN","BRANCH_ADMIN","STUDENT"))
	api.PUT("/subjects/:subjectId", subjectHandler.UpdateSubject, authMiddleware, auth.RequireRoles("SUPER_ADMIN"))
	api.DELETE("/subjects/:subjectId", subjectHandler.DeleteSubject, authMiddleware, auth.RequireRoles("SUPER_ADMIN"))
}