package producthdl

import "gin-swagger-api/internal/domain"

// ProductResponse represents the API response for a product
type ProductResponse struct {
	ID          string  `json:"id" example:"1"`
	Name        string  `json:"name" example:"Laptop"`
	Description string  `json:"description" example:"Gaming laptop"`
	Price       float64 `json:"price" example:"25000.50"`
	Stock       int     `json:"stock" example:"10"`
}

// CreateProductRequest represents the request body for creating a product
type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required" example:"Laptop"`
	Description string  `json:"description" example:"Gaming laptop"`
	Price       float64 `json:"price" binding:"required,gt=0" example:"25000.50"`
	Stock       int     `json:"stock" binding:"gte=0" example:"10"`
}

// UpdateProductRequest represents the request body for updating a product
type UpdateProductRequest struct {
	Name        string  `json:"name" example:"Laptop"`
	Description string  `json:"description" example:"Gaming laptop"`
	Price       float64 `json:"price" example:"25000.50"`
	Stock       int     `json:"stock" example:"10"`
}

// toProductResponse converts domain.Product to ProductResponse
func toProductResponse(product domain.Product) ProductResponse {
	return ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}
}
