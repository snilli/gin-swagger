package userrepo

import (
	"context"
	"strconv"

	"gin-swagger-api/internal/domain"
	portuserrepo "gin-swagger-api/internal/port/repository/userrepo"

	ormprovider "github.com/example/orm-provider-api"
)

// Repository implements the user repository interface
type Repository struct {
	db *ormprovider.Client
}

// New creates a new user repository
func New(db *ormprovider.Client) portuserrepo.Repository {
	return &Repository{db: db}
}

// GetAll retrieves all users
func (r *Repository) GetAll(ctx context.Context) ([]domain.User, error) {
	entUsers, err := r.db.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	users := make([]domain.User, len(entUsers))
	for i, entUser := range entUsers {
		users[i] = domain.User{
			ID:    strconv.Itoa(entUser.ID),
			Name:  entUser.Name,
			Email: entUser.Email,
		}
	}
	return users, nil
}

// GetByID retrieves a user by ID
func (r *Repository) GetByID(ctx context.Context, id int) (*domain.User, error) {
	entUser, err := r.db.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:    strconv.Itoa(entUser.ID),
		Name:  entUser.Name,
		Email: entUser.Email,
	}, nil
}

// Create creates a new user
func (r *Repository) Create(ctx context.Context, name, email string) (*domain.User, error) {
	entUser, err := r.db.User.Create().
		SetName(name).
		SetEmail(email).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:    strconv.Itoa(entUser.ID),
		Name:  entUser.Name,
		Email: entUser.Email,
	}, nil
}

// Update updates a user
func (r *Repository) Update(ctx context.Context, id int, name, email string) (*domain.User, error) {
	entUser, err := r.db.User.UpdateOneID(id).
		SetName(name).
		SetEmail(email).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:    strconv.Itoa(entUser.ID),
		Name:  entUser.Name,
		Email: entUser.Email,
	}, nil
}

// Delete deletes a user
func (r *Repository) Delete(ctx context.Context, id int) error {
	return r.db.User.DeleteOneID(id).Exec(ctx)
}
