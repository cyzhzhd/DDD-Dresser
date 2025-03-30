package products

import (
	"dresser/internal/domain/categoriess"
	"dresser/internal/domain/prices"
	dproducts "dresser/internal/domain/products"
	"dresser/internal/domain/sizes"
)

// ProductDTO represents the application's view of a product
type ProductDTO struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	BrandID  int      `json:"brand_id"`
	Category string   `json:"category"`
	Price    PriceDTO `json:"price"`
	Size     SizeDTO  `json:"size"`
}

// PriceDTO represents the price information in the DTO
type PriceDTO struct {
	Amount       int    `json:"amount"`
	Currency     string `json:"currency"`
	DisplayValue string `json:"display_value"`
}

// SizeDTO represents the size information in the DTO
type SizeDTO struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

func product2DTO(product *dproducts.Product) *ProductDTO {
	price := product.GetPrice()
	size := product.GetSize()

	return &ProductDTO{
		ID:       product.GetID(),
		Name:     product.GetName(),
		BrandID:  int(product.GetBrandID()),
		Category: product.GetCategory(),
		Price: PriceDTO{
			Amount:       price.Amount(),
			Currency:     string(price.Currency()),
			DisplayValue: price.String(),
		},
		Size: SizeDTO{
			Value: size.Value(),
			Type:  string(size.Type()),
		},
	}
}

// This is a utility function to help transform DTO objects back to domain objects
func createPriceFromDTO(amount int, currency string) (*prices.Price, error) {
	currencyEnum := prices.Currency(currency)
	return prices.NewPrice(amount, currencyEnum)
}

func createSizeFromDTO(value string, category string) (*sizes.Size, error) {
	categoryEnum := categoriess.Category(category)
	return sizes.NewSize(value, categoryEnum)
}
