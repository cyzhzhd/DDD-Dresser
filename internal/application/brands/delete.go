package brands

import (
	"context"
	dbrands "dresser/internal/domain/brands"
)

type IBrandDeleteApplicationService interface {
	Delete(ctx context.Context, id int) error
}
type BrandDeleteService struct {
	repository dbrands.Repository
}

func NewBrandDeleteService(repository dbrands.Repository) *BrandDeleteService {
	return &BrandDeleteService{
		repository: repository,
	}
}

func (s *BrandDeleteService) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, dbrands.ID(id))
	if err != nil {
		return err
	}
	return nil
}
