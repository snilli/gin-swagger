package producthdl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProducts godoc
// @Summary List all products
// @Description Get a list of all products
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {array} ProductResponse
// @Failure 500 {object} ErrorResponse
// @Router /products [get]
func (h *Handler) GetProducts(c *gin.Context) {
	products, err := h.productService.GetProducts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	response := make([]ProductResponse, len(products))
	for i, product := range products {
		response[i] = toProductResponse(product)
	}

	c.JSON(http.StatusOK, response)
}
