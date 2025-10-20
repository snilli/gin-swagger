package orderrepo

import (
	"context"
	"strconv"

	"gin-swagger-api/internal/domain"
	portorderrepo "gin-swagger-api/internal/port/repository/orderrepo"

	ormprovider "github.com/example/orm-provider-api"
)

// Repository implements the order repository interface
type Repository struct {
	db *ormprovider.Client
}

// New creates a new order repository
func New(db *ormprovider.Client) portorderrepo.Repository {
	return &Repository{db: db}
}

// GetAll retrieves all orders
func (r *Repository) GetAll(ctx context.Context) ([]domain.Order, error) {
	entOrders, err := r.db.Order.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	orders := make([]domain.Order, len(entOrders))
	for i, entOrder := range entOrders {
		orders[i] = domain.Order{
			ID:         strconv.Itoa(entOrder.ID),
			UserID:     entOrder.UserID,
			ProductID:  entOrder.ProductID,
			Quantity:   entOrder.Quantity,
			TotalPrice: entOrder.TotalPrice,
			Status:     entOrder.Status,
		}
	}
	return orders, nil
}

// GetByID retrieves an order by ID
func (r *Repository) GetByID(ctx context.Context, id int) (*domain.Order, error) {
	entOrder, err := r.db.Order.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &domain.Order{
		ID:         strconv.Itoa(entOrder.ID),
		UserID:     entOrder.UserID,
		ProductID:  entOrder.ProductID,
		Quantity:   entOrder.Quantity,
		TotalPrice: entOrder.TotalPrice,
		Status:     entOrder.Status,
	}, nil
}

// Create creates a new order
func (r *Repository) Create(ctx context.Context, userID, productID, quantity int, totalPrice float64, status string) (*domain.Order, error) {
	entOrder, err := r.db.Order.Create().
		SetUserID(userID).
		SetProductID(productID).
		SetQuantity(quantity).
		SetTotalPrice(totalPrice).
		SetStatus(status).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &domain.Order{
		ID:         strconv.Itoa(entOrder.ID),
		UserID:     entOrder.UserID,
		ProductID:  entOrder.ProductID,
		Quantity:   entOrder.Quantity,
		TotalPrice: entOrder.TotalPrice,
		Status:     entOrder.Status,
	}, nil
}

// Update updates an order
func (r *Repository) Update(ctx context.Context, id, quantity int, totalPrice float64, status string) (*domain.Order, error) {
	entOrder, err := r.db.Order.UpdateOneID(id).
		SetQuantity(quantity).
		SetTotalPrice(totalPrice).
		SetStatus(status).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &domain.Order{
		ID:         strconv.Itoa(entOrder.ID),
		UserID:     entOrder.UserID,
		ProductID:  entOrder.ProductID,
		Quantity:   entOrder.Quantity,
		TotalPrice: entOrder.TotalPrice,
		Status:     entOrder.Status,
	}, nil
}

// Delete deletes an order
func (r *Repository) Delete(ctx context.Context, id int) error {
	return r.db.Order.DeleteOneID(id).Exec(ctx)
}
