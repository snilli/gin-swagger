package orderhdl

import (
	"gin-swagger-api/internal/middleware"
	"gin-swagger-api/internal/port/service/ordersvc"

	"github.com/gin-gonic/gin"
)

// Handler handles order-related HTTP requests
type Handler struct {
	orderService ordersvc.Service
}

// NewHandler creates a new order handler
func NewHandler(orderService ordersvc.Service) *Handler {
	return &Handler{
		orderService: orderService,
	}
}

// RegisterRoutes registers all order routes
func (h *Handler) RegisterRoutes(rg *gin.RouterGroup) {
	orders := rg.Group("/orders")
	orders.Use(middleware.Logger())     // Apply logger to all order routes
	orders.Use(middleware.Auth())       // Apply auth to all order routes (example)
	{
		orders.POST("", h.CreateOrder)
		orders.GET("/:id", h.GetOrder)
		orders.PUT("/:id", h.UpdateOrder)
		orders.DELETE("/:id", h.DeleteOrder)
		orders.GET("", h.GetOrders)
	}
}
