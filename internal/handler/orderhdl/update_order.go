package orderhdl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateOrder godoc
// @Summary Update an order
// @Description Update an order's information by ID
// @Tags orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Param order body UpdateOrderRequest true "Order information"
// @Success 200 {object} OrderResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /orders/{id} [put]
func (h *Handler) UpdateOrder(c *gin.Context) {
	id := c.Param("id")

	var req UpdateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	order, err := h.orderService.UpdateOrder(
		c.Request.Context(),
		id,
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

	c.JSON(http.StatusOK, toOrderResponse(*order))
}
