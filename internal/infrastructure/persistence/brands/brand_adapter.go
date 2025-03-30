package brands

import (
	appBrands "dresser/internal/application/brands"
	"dresser/internal/domain/brands"
)

// BrandAdapter converts between domain and application brand models
type BrandAdapter struct{}

func NewBrandAdapter() *BrandAdapter {
	return &BrandAdapter{}
}

// ToDomain converts application brand to domain brand
func (a *BrandAdapter) ToDomain(appBrand *appBrands.BrandDTO) (*brands.Brand, error) {
	return brands.New(appBrand.ID, appBrand.Name)
}

// ToApplication converts domain brand to application brand
func (a *BrandAdapter) ToApplication(domainBrand *brands.Brand) *appBrands.BrandDTO {
	return &appBrands.BrandDTO{
		ID:   int(domainBrand.ID),
		Name: domainBrand.Name,
	}
}
