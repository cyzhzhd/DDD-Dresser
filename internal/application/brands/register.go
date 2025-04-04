package brands

import (
	"context"
	dbrands "dresser/internal/domain/brands"
)

type IBrandRegisterApplicationService interface {
	Register(ctx context.Context, cmd RegisterBrandCommand) (*BrandDTO, error)
}

type BrandRegisterService struct {
	repository dbrands.Repository
	factory    dbrands.Factory
}

func NewBrandRegisterService(repository dbrands.Repository, factory dbrands.Factory) *BrandRegisterService {
	return &BrandRegisterService{
		repository: repository,
		factory:    factory,
	}
}

type RegisterBrandCommand struct {
	Name string
}

func (s *BrandRegisterService) Register(ctx context.Context, cmd RegisterBrandCommand) (*BrandDTO, error) {
	brand, err := s.factory.CreateBrand(ctx, cmd.Name)
	if err != nil {
		return nil, err
	}

	// Save brand
	res, err := s.repository.Create(ctx, brand)
	if err != nil {
		return nil, err
	}

	return brand2DTO(res), nil
}
