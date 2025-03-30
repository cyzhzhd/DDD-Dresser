package products

import (
	"context"
	dproducts "dresser/internal/domain/products"
)

type IProductQueryApplicationService interface {
	GetAll(ctx context.Context) ([]*ProductDTO, error)
	Get(ctx context.Context, id int) (*ProductDTO, error)
	GetByBrand(ctx context.Context, brandID int) ([]*ProductDTO, error)
	GetByCategory(ctx context.Context, category string) ([]*ProductDTO, error)
}

type ProductQueryService struct {
	repository dproducts.Repository
}

func NewProductQueryService(repository dproducts.Repository) *ProductQueryService {
	return &ProductQueryService{
		repository: repository,
	}
}

func (s *ProductQueryService) GetAll(ctx context.Context) ([]*ProductDTO, error) {
	products, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*ProductDTO, 0, len(products))
	for _, product := range products {
		result = append(result, product2DTO(product))
	}
	return result, nil
}

func (s *ProductQueryService) Get(ctx context.Context, id int) (*ProductDTO, error) {
	product, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return product2DTO(product), nil
}

func (s *ProductQueryService) GetByBrand(ctx context.Context, brandID int) ([]*ProductDTO, error) {
	products, err := s.repository.FindByBrand(ctx, brandID)
	if err != nil {
		return nil, err
	}

	result := make([]*ProductDTO, 0, len(products))
	for _, product := range products {
		result = append(result, product2DTO(product))
	}
	return result, nil
}

func (s *ProductQueryService) GetByCategory(ctx context.Context, category string) ([]*ProductDTO, error) {
	products, err := s.repository.FindByCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	result := make([]*ProductDTO, 0, len(products))
	for _, product := range products {
		result = append(result, product2DTO(product))
	}
	return result, nil
}
