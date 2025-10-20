package usersvc

import (
	port "gin-swagger-api/internal/port/service/usersvc"
	userrepo "gin-swagger-api/internal/port/repository/userrepo"
)

// Service implements port.Service interface
type Service struct {
	userRepo userrepo.Repository
}

// New creates a new user service with user repository
func New(userRepo userrepo.Repository) port.Service {
	return &Service{
		userRepo: userRepo,
	}
}
