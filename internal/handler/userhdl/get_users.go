package userhdl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers godoc
// @Summary Get all users
// @Description Get all users from the system
// @Tags users
// @Produce json
// @Success 200 {array} UserResponse
// @Failure 500 {object} ErrorResponse
// @Router /users [get]
func (h *Handler) GetUsers(c *gin.Context) {
	users, err := h.userService.GetUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	response := make([]UserResponse, 0, len(users))
	for _, user := range users {
		response = append(response, toUserResponse(user))
	}

	c.JSON(http.StatusOK, response)
}
