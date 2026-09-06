package subjects

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoute(e *echo.Echo, db *gorm.DB) {
	subjectRepo := NewRepository(db)
	subjectService := NewService(subjectRepo)
	subjectHandler := NewHandler(subjectService)

	subjects := e.Group("/api/v1")

	subjects.POST("/subjects", subjectHandler.CreateSubject)
	subjects.GET("/subjects", subjectHandler.GetAll)
	subjects.GET("/subjects/:subjectId", subjectHandler.FindById)
	subjects.GET("/query", subjectHandler.CheckQuery)
	subjects.PUT("/subjects/:subjectId", subjectHandler.UpdateSubject)
	subjects.DELETE("/subjects/:subjectId", subjectHandler.DeleteSubject)
}