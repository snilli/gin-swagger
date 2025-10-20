package orderhdl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order with the provided information
// @Tags orders
// @Accept json
// @Produce json
// @Param order body CreateOrderRequest true "Order information"
// @Success 201 {object} OrderResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /orders [post]
func (h *Handler) CreateOrder(c *gin.Context) {
	var req CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	order, err := h.orderService.CreateOrder(
		c.Request.Context(),
		req.UserID,
		req.ProductID,
		req.Quantity,
		req.TotalPrice,
		req.Status,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, toOrderResponse(*order))
}
