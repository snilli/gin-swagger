package userhdl

import (
	"gin-swagger-api/internal/middleware"
	"gin-swagger-api/internal/port/service/usersvc"

	"github.com/gin-gonic/gin"
)

// Handler handles user-related HTTP requests
type Handler struct {
	userService usersvc.Service
}

// NewHandler creates a new user handler
func NewHandler(userService usersvc.Service) *Handler {
	return &Handler{
		userService: userService,
	}
}

// RegisterRoutes registers all user routes
func (h *Handler) RegisterRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	users.Use(middleware.Logger()) // Apply logger to all user routes
	{
		users.POST("", h.CreateUser)
		users.GET("/:id", h.GetUser)
		users.PUT("/:id", h.UpdateUser)
		users.DELETE("/:id", h.DeleteUser)
		users.GET("", h.GetUsers)
	}
}
