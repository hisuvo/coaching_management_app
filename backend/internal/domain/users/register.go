package users

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoute(e *echo.Echo, db *gorm.DB) {

	userRepo := NewRepository(db)
	userService := NewService(userRepo)
	userHandler := NewHandler(userService)

	users := e.Group("/api/v1")

	users.POST("/auth/register",userHandler.Register)

	users.GET("/users/:email", userHandler.FindByEmail)

	// users.GET("/users/:id", func(c *echo.Context) error {
	// 	return c.JSON(200, map[string]any{
	// 		"id": 1,
	// 		"name": "Leanne Graham",
	// 		"username": "Bret",
	// 		"email": "Sincere@april.biz",
	// 		"isActive": true,
	// 		"roles": []string{"User", "Editor"},
	// 		"address": map[string]string{
	// 		"street": "Kulas Light",
	// 		"suite": "Apt. 556",
	// 		"city": "Gwenborough",
	// 		"zipcode": "92998-3874",
	// 		},
	// 		"phone": "1-770-736-8031 x56442",
	// 		"website": "hildegard.org",
	// 	})
	// })
}