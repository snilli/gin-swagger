package orderhdl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetOrder godoc
// @Summary Get order by ID
// @Description Get an order by its ID
// @Tags orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} OrderResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /orders/{id} [get]
func (h *Handler) GetOrder(c *gin.Context) {
	id := c.Param("id")

	order, err := h.orderService.GetOrder(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "order not found"})
		return
	}

	c.JSON(http.StatusOK, toOrderResponse(*order))
}
