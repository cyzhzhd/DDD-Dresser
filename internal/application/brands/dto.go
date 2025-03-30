package brands

import (
	dbrands "dresser/internal/domain/brands"
)

// BrandDTO represents the application's view of a brand
type BrandDTO struct {
	ID         int
	Name       string
	Categories []string
}

func brand2DTO(brand *dbrands.Brand) *BrandDTO {
	return &BrandDTO{
		ID:   int(brand.GetID()),
		Name: brand.GetName(),
	}
}
