package main

import (
	"log"
	"user-management-api/controllers"
	_ "user-management-api/docs"
	"user-management-api/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title User Management API
// @version 1.0
// @description This is a user management API server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	// Database connection setup
	dataSourceName := "postgres://kpk:qwerty@localhost:5432/user_management_db"
	controllers.InitDB(dataSourceName)

	// Echo instance setup
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Enable CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4200"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	// Initialize routes
	routes.Init(e)

	// Swagger endpoint
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	log.Println("Starting server on :8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}
