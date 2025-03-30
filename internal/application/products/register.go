package products

import (
	"context"
	dproducts "dresser/internal/domain/products"
)

type IProductRegisterApplicationService interface {
	Register(ctx context.Context, cmd RegisterProductCommand) (*ProductDTO, error)
}

type ProductRegisterService struct {
	repository dproducts.Repository
	factory    dproducts.Factory
}

func NewProductRegisterService(repository dproducts.Repository, factory dproducts.Factory) *ProductRegisterService {
	return &ProductRegisterService{
		repository: repository,
		factory:    factory,
	}
}

type RegisterProductCommand struct {
	Name     string `json:"name"`
	BrandID  int    `json:"brand_id"`
	Category string `json:"category"`
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
	Size     string `json:"size"`
}

func (s *ProductRegisterService) Register(ctx context.Context, cmd RegisterProductCommand) (*ProductDTO, error) {
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

	// Create new product using factory
	product, err := s.factory.CreateProduct(ctx, cmd.Name, cmd.BrandID, cmd.Category, price, size)
	if err != nil {
		return nil, err
	}

	// Save product
	res, err := s.repository.Create(ctx, product)
	if err != nil {
		return nil, err
	}

	return product2DTO(res), nil
}
