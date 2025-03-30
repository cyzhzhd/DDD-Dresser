package brands

import (
	"context"
)

// Repository defines the interface for brand persistence
type Repository interface {
	// NextID generates the next sequential ID
	NextID() ID

	// Create saves a new brand
	Create(ctx context.Context, brand *Brand) (*Brand, error)

	// Update updates an existing brand
	Update(ctx context.Context, brand *Brand) error

	// Delete removes a brand by ID
	Delete(ctx context.Context, id ID) error

	// FindByID retrieves a brand by ID
	FindByID(ctx context.Context, id ID) (*Brand, error)

	// FindAll retrieves all brands
	FindAll(ctx context.Context) ([]*Brand, error)

	// FindByCategory retrieves all brands that have the specified category
	FindByCategory(ctx context.Context, category string) ([]*Brand, error)

	// Exists checks if a brand exists by ID
	Exists(ctx context.Context, id ID) (bool, error)
}
