package brands

import "context"

type BrandService struct {
	brandRepository Repository
}

func NewBrandService(brandRepository Repository) *BrandService {
	return &BrandService{brandRepository: brandRepository}
}

func (s *BrandService) RegisterBrand(ctx context.Context, brand *Brand) (*Brand, error) {
	return s.brandRepository.Create(ctx, brand)
}
