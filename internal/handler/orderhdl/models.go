package orderhdl

import "gin-swagger-api/internal/domain"

// OrderResponse represents the API response for an order
type OrderResponse struct {
	ID         string  `json:"id" example:"1"`
	UserID     int     `json:"user_id" example:"1"`
	ProductID  int     `json:"product_id" example:"1"`
	Quantity   int     `json:"quantity" example:"2"`
	TotalPrice float64 `json:"total_price" example:"50000.00"`
	Status     string  `json:"status" example:"pending"`
}

// CreateOrderRequest represents the request body for creating an order
type CreateOrderRequest struct {
	UserID     int     `json:"user_id" binding:"required" example:"1"`
	ProductID  int     `json:"product_id" binding:"required" example:"1"`
	Quantity   int     `json:"quantity" binding:"required,gt=0" example:"2"`
	TotalPrice float64 `json:"total_price" binding:"required,gt=0" example:"50000.00"`
	Status     string  `json:"status" example:"pending"`
}

// UpdateOrderRequest represents the request body for updating an order
type UpdateOrderRequest struct {
	UserID     int     `json:"user_id" example:"1"`
	ProductID  int     `json:"product_id" example:"1"`
	Quantity   int     `json:"quantity" example:"2"`
	TotalPrice float64 `json:"total_price" example:"50000.00"`
	Status     string  `json:"status" example:"completed"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error" example:"error message"`
}

// toOrderResponse converts domain.Order to OrderResponse
func toOrderResponse(order domain.Order) OrderResponse {
	return OrderResponse{
		ID:         order.ID,
		UserID:     order.UserID,
		ProductID:  order.ProductID,
		Quantity:   order.Quantity,
		TotalPrice: order.TotalPrice,
		Status:     order.Status,
	}
}
