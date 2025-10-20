package ordersvc

import (
	port "gin-swagger-api/internal/port/service/ordersvc"
	orderrepo "gin-swagger-api/internal/port/repository/orderrepo"
)

// Service implements port.Service interface
type Service struct {
	orderRepo orderrepo.Repository
}

// New creates a new order service with order repository
func New(orderRepo orderrepo.Repository) port.Service {
	return &Service{
		orderRepo: orderRepo,
	}
}
