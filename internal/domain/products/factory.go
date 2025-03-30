package products

import (
	"context"
	"dresser/internal/domain/brands"
	"dresser/internal/domain/prices"
	"dresser/internal/domain/sizes"
	"fmt"
)

type Factory interface {
	CreateProduct(ctx context.Context, name string, brandID int, category string, price *prices.Price, size *sizes.Size) (*Product, error)
}

type ProductFactory struct {
	repository      Repository
	brandRepository brands.Repository
	nextIDFunc      func() int
}

func NewProductFactory(repository Repository, brandRepository brands.Repository) *ProductFactory {
	return &ProductFactory{
		repository:      repository,
		brandRepository: brandRepository,
		nextIDFunc: func() int {
			return repository.NextID()
		},
	}
}

// CreateProduct creates a new valid product
func (f *ProductFactory) CreateProduct(ctx context.Context, name string, brandID int, category string, price *prices.Price, size *sizes.Size) (*Product, error) {
	// Validate name
	if name == "" {
		return nil, fmt.Errorf("product name cannot be empty")
	}

	// Verify brand exists
	brandExists, err := f.brandRepository.Exists(ctx, brands.ID(brandID))
	if err != nil {
		return nil, err
	}
	if !brandExists {
		return nil, fmt.Errorf("brand with ID %d does not exist", brandID)
	}

	product := New(
		f.nextIDFunc(),
		name,
		brands.ID(brandID),
		category,
		*price,
		*size,
	)

	return product, nil
}
