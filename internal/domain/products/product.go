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

// New creates a new Product
func New(id int, name string, brandID brands.ID, category string, price prices.Price, size sizes.Size) *Product {
	return &Product{
		ID:       id,
		Name:     name,
		Brand:    brandID,
		Category: category,
		Price:    price,
		Size:     size,
	}
}

// GetID returns the product ID
func (p *Product) GetID() int {
	return p.ID
}

// GetName returns the product name
func (p *Product) GetName() string {
	return p.Name
}

// GetBrandID returns the brand ID
func (p *Product) GetBrandID() brands.ID {
	return p.Brand
}

// GetCategory returns the product category
func (p *Product) GetCategory() string {
	return p.Category
}

// GetPrice returns the product price
func (p *Product) GetPrice() prices.Price {
	return p.Price
}

// GetSize returns the product size
func (p *Product) GetSize() sizes.Size {
	return p.Size
}
