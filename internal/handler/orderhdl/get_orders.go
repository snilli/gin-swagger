package orderhdl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetOrders godoc
// @Summary List all orders
// @Description Get a list of all orders
// @Tags orders
// @Accept json
// @Produce json
// @Success 200 {array} OrderResponse
// @Failure 500 {object} ErrorResponse
// @Router /orders [get]
func (h *Handler) GetOrders(c *gin.Context) {
	orders, err := h.orderService.GetOrders(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	response := make([]OrderResponse, len(orders))
	for i, order := range orders {
		response[i] = toOrderResponse(order)
	}

	c.JSON(http.StatusOK, response)
}
