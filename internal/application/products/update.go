package products

import (
	"context"
	"dresser/internal/domain/brands"
	dproducts "dresser/internal/domain/products"
	"errors"
	"fmt"
)

type IProductUpdateApplicationService interface {
	Update(ctx context.Context, cmd UpdateProductCommand) (*ProductDTO, error)
}

type ProductUpdateService struct {
	repository dproducts.Repository
	factory    dproducts.Factory
}

func NewProductUpdateService(repository dproducts.Repository, factory dproducts.Factory) *ProductUpdateService {
	return &ProductUpdateService{
		repository: repository,
		factory:    factory,
	}
}

type UpdateProductCommand struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	BrandID  int    `json:"brand_id"`
	Category string `json:"category"`
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
	Size     string `json:"size"`
}

func (s *ProductUpdateService) Update(ctx context.Context, cmd UpdateProductCommand) (*ProductDTO, error) {
	// Check if product exists
	existingProduct, err := s.repository.FindByID(ctx, cmd.ID)
	if err != nil {
		return nil, err
	}
	if existingProduct == nil {
		return nil, errors.New("product not found")
	}

	// Create price object
	price, err := createPriceFromDTO(cmd.Amount, cmd.Currency)
	if err != nil {
		return nil, err
	}

	// Create size object
	size, err := createSizeFromDTO(cmd.Size, cmd.Category)
	if err != nil {
		return nil, err
	}

	// Create updated product
	updatedProduct := dproducts.New(
		cmd.ID,
		cmd.Name,
		brands.ID(cmd.BrandID),
		cmd.Category,
		*price,
		*size,
	)

	// Update product
	err = s.repository.Update(ctx, updatedProduct)
	if err != nil {
		return nil, fmt.Errorf("failed to update product: %w", err)
	}

	// Get updated product
	result, err := s.repository.FindByID(ctx, cmd.ID)
	if err != nil {
		return nil, err
	}

	return product2DTO(result), nil
}
