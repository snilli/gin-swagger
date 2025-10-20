package productrepo

import (
	"context"
	"strconv"

	"gin-swagger-api/internal/domain"
	portproductrepo "gin-swagger-api/internal/port/repository/productrepo"

	ormprovider "github.com/example/orm-provider-api"
)

// Repository implements the product repository interface
type Repository struct {
	db *ormprovider.Client
}

// New creates a new product repository
func New(db *ormprovider.Client) portproductrepo.Repository {
	return &Repository{db: db}
}

// GetAll retrieves all products
func (r *Repository) GetAll(ctx context.Context) ([]domain.Product, error) {
	entProducts, err := r.db.Product.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	products := make([]domain.Product, len(entProducts))
	for i, entProduct := range entProducts {
		products[i] = domain.Product{
			ID:          strconv.Itoa(entProduct.ID),
			Name:        entProduct.Name,
			Description: entProduct.Description,
			Price:       entProduct.Price,
			Stock:       entProduct.Stock,
		}
	}
	return products, nil
}

// GetByID retrieves a product by ID
func (r *Repository) GetByID(ctx context.Context, id int) (*domain.Product, error) {
	entProduct, err := r.db.Product.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &domain.Product{
		ID:          strconv.Itoa(entProduct.ID),
		Name:        entProduct.Name,
		Description: entProduct.Description,
		Price:       entProduct.Price,
		Stock:       entProduct.Stock,
	}, nil
}

// Create creates a new product
func (r *Repository) Create(ctx context.Context, name, description string, price float64, stock int) (*domain.Product, error) {
	entProduct, err := r.db.Product.Create().
		SetName(name).
		SetDescription(description).
		SetPrice(price).
		SetStock(stock).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &domain.Product{
		ID:          strconv.Itoa(entProduct.ID),
		Name:        entProduct.Name,
		Description: entProduct.Description,
		Price:       entProduct.Price,
		Stock:       entProduct.Stock,
	}, nil
}

// Update updates a product
func (r *Repository) Update(ctx context.Context, id int, name, description string, price float64, stock int) (*domain.Product, error) {
	entProduct, err := r.db.Product.UpdateOneID(id).
		SetName(name).
		SetDescription(description).
		SetPrice(price).
		SetStock(stock).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &domain.Product{
		ID:          strconv.Itoa(entProduct.ID),
		Name:        entProduct.Name,
		Description: entProduct.Description,
		Price:       entProduct.Price,
		Stock:       entProduct.Stock,
	}, nil
}

// Delete deletes a product
func (r *Repository) Delete(ctx context.Context, id int) error {
	return r.db.Product.DeleteOneID(id).Exec(ctx)
}
