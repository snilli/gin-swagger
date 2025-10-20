package producthdl

import (
	"gin-swagger-api/internal/middleware"
	"gin-swagger-api/internal/port/service/productsvc"

	"github.com/gin-gonic/gin"
)

// Handler handles product-related HTTP requests
type Handler struct {
	productService productsvc.Service
}

// NewHandler creates a new product handler
func NewHandler(productService productsvc.Service) *Handler {
	return &Handler{
		productService: productService,
	}
}

// RegisterRoutes registers all product routes
func (h *Handler) RegisterRoutes(rg *gin.RouterGroup) {
	products := rg.Group("/products")
	products.Use(middleware.Logger()) // Apply logger to all product routes
	{
		products.POST("", h.CreateProduct)
		products.GET("/:id", h.GetProduct)
		products.PUT("/:id", h.UpdateProduct)
		products.DELETE("/:id", h.DeleteProduct)
		products.GET("", h.GetProducts)
	}
}
