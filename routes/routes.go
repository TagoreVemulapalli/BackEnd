package routes

import (
	"user-management-api/controllers"

	"github.com/labstack/echo/v4"
)

// Init initializes the routes
func Init(e *echo.Echo) {
	// User routes
	e.GET("/api/users", controllers.GetUsers)
	e.GET("/api/users/:user_id", controllers.GetUserById)
	e.POST("/api/users", controllers.CreateUser)
	e.PUT("/api/users/:user_id", controllers.UpdateUser)
	e.DELETE("/api/users/:user_id", controllers.DeleteUser)
}
