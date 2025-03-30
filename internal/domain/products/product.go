package products

import (
	"dresser/internal/domain/brands"
	"dresser/internal/domain/prices"
	"dresser/internal/domain/sizes"
)

type (
	Product struct {
		ID       int
		Name     string
		Brand    brands.ID
		Category string
		Price    prices.Price
		Size     sizes.Size
	}
)
