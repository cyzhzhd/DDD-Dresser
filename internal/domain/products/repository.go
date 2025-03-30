package products

import (
	"context"
)

// Repository defines the interface for product persistence
type Repository interface {
	// NextID generates the next sequential ID
	NextID() int

	// Create saves a new product
	Create(ctx context.Context, product *Product) (*Product, error)

	// Update updates an existing product
	Update(ctx context.Context, product *Product) error

	// Delete removes a product by ID
	Delete(ctx context.Context, id int) error

	// FindByID retrieves a product by ID
	FindByID(ctx context.Context, id int) (*Product, error)

	// FindAll retrieves all products
	FindAll(ctx context.Context) ([]*Product, error)

	// FindByBrand retrieves all products for a specific brand
	FindByBrand(ctx context.Context, brandID int) ([]*Product, error)

	// FindByCategory retrieves all products that have the specified category
	FindByCategory(ctx context.Context, category string) ([]*Product, error)

	// Exists checks if a product exists by ID
	Exists(ctx context.Context, id int) (bool, error)
}
