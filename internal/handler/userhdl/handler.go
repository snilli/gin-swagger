package userhdl

import "meek/internal/port/usersvc"

// Handler handles user-related HTTP requests
type Handler struct {
	userService usersvc.Service
}

// NewHandler creates a new user handler
func NewHandler(userService usersvc.Service) *Handler {
	return &Handler{
		userService: userService,
	}
}
