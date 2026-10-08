package routes

import (
	"coaching_backend/internal/config"
	"coaching_backend/internal/domain/assignments"
	"coaching_backend/internal/domain/auth"
	"coaching_backend/internal/domain/branches"
	"coaching_backend/internal/domain/coaching"
	coachingsubject "coaching_backend/internal/domain/coachingSubject"
	"coaching_backend/internal/domain/students"
	"coaching_backend/internal/domain/subjects"
	"coaching_backend/internal/domain/submissions"
	"coaching_backend/internal/domain/teachers"
	"coaching_backend/internal/domain/users"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, cnfg *config.Config){
	
	userRepo := users.NewRepository(db)
    authRepo := auth.NewRepository(db)

    tokenManager := auth.NewTokenManager(
        cnfg.JWT_ACCESS_SECRET,
        "coaching-management-api",
        cnfg.JWT_ACCESS_EXPIRES_IN,
        cnfg.JWT_REFRESH_EXPIRES_IN,
    )

    authService := auth.NewService(userRepo, authRepo, tokenManager)
    authHandler := auth.NewHandler(authService, authRepo, *cnfg, *tokenManager)

	authMiddleware := authHandler.AuthMiddleware;

	api := e.Group("api/v1")

	// routes
	auth.RegisterRoute(e, db, cnfg)
	coaching.RegisterRoute(api, db, authMiddleware)
	branches.RegisterRoute(api, db, authMiddleware)
	users.RegisterRoute(e, db)
	subjects.RegisterRoute(api, db, authMiddleware)
	teachers.RegisterRoute(api, db, authMiddleware)
	coachingsubject.RegisterRoute(api, db, authMiddleware)
	students.RegisterRoute(e, db)
	submissions.RegisterRoute(e, db)
	assignments.RegisterRoute(e, db)

}