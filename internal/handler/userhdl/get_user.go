package userhdl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUser godoc
// @Summary Get user by ID
// @Description Get a single user by ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} UserResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/{id} [get]
func (h *Handler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.userService.GetUser(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, toUserResponse(*user))
}
