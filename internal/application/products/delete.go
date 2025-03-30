package products

import (
	"context"
	dproducts "dresser/internal/domain/products"
)

type IProductDeleteApplicationService interface {
	Delete(ctx context.Context, id int) error
}

type ProductDeleteService struct {
	repository dproducts.Repository
}

func NewProductDeleteService(repository dproducts.Repository) *ProductDeleteService {
	return &ProductDeleteService{
		repository: repository,
	}
}

func (s *ProductDeleteService) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
