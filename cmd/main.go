package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "meek/docs"
	"meek/internal/handler"
	"meek/internal/handler/userhdl"
	"meek/internal/service/usersvc"
)

// @title Meek API
// @version 1.0
// @description API documentation for Meek service
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @schemes http https
func main() {
	r := gin.Default()

	// Initialize services
	userService := usersvc.New()

	// Initialize handlers
	healthHandler := handler.NewHealthHandler()
	userHandler := userhdl.NewHandler(userService)

	// Health check endpoint
	r.GET("/health", healthHandler.HealthCheck)

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", userHandler.GetUsers)
		v1.GET("/users/:id", userHandler.GetUser)
		v1.POST("/users", userHandler.CreateUser)
		v1.PUT("/users/:id", userHandler.UpdateUser)
		v1.DELETE("/users/:id", userHandler.DeleteUser)
	}

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
