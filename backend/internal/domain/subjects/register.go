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

	subjects := api.Group("/subject")
	subjects.Use(authMiddleware)

	
	subjects.POST("", subjectHandler.CreateSubject)
	subjects.GET("", subjectHandler.GetAll)

	member := subjects.Group("")
	member.Use(auth.RequireRoles("MEMBER","STUDENT"))
	member.GET("/:subjectId", subjectHandler.FindById)
	
	subjects.GET("/query", subjectHandler.CheckQuery)
	subjects.PUT("/:subjectId", subjectHandler.UpdateSubject)
	subjects.DELETE("/:subjectId", subjectHandler.DeleteSubject)
}