package producthdl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProduct godoc
// @Summary Get product by ID
// @Description Get a product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} ProductResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /products/{id} [get]
func (h *Handler) GetProduct(c *gin.Context) {
	id := c.Param("id")

	product, err := h.productService.GetProduct(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "product not found"})
		return
	}

	c.JSON(http.StatusOK, toProductResponse(*product))
}
