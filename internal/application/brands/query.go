package brands

import (
	"context"
	dbrands "dresser/internal/domain/brands"
)

type IBrandQueryApplicationService interface {
	GetAll(ctx context.Context) ([]*BrandDTO, error)
	Get(ctx context.Context, id int) (*BrandDTO, error)
}

type BrandQueryService struct {
	repository dbrands.Repository
}

func NewBrandQueryService(repository dbrands.Repository) *BrandQueryService {
	return &BrandQueryService{
		repository: repository,
	}
}

func (s *BrandQueryService) GetAll(ctx context.Context) ([]*BrandDTO, error) {
	brands, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*BrandDTO, 0, len(brands))
	for _, brand := range brands {
		result = append(result, brand2DTO(brand))
	}
	return result, nil
}

func (s *BrandQueryService) Get(ctx context.Context, id int) (*BrandDTO, error) {
	brand, err := s.repository.FindByID(ctx, dbrands.ID(id))
	if err != nil {
		return nil, err
	}
	return brand2DTO(brand), nil
}
