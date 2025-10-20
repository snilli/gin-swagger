package productsvc

import (
	port "gin-swagger-api/internal/port/service/productsvc"
	productrepo "gin-swagger-api/internal/port/repository/productrepo"
)

// Service implements port.Service interface
type Service struct {
	productRepo productrepo.Repository
}

// New creates a new product service with product repository
func New(productRepo productrepo.Repository) port.Service {
	return &Service{
		productRepo: productRepo,
	}
}
