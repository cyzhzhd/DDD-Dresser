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

func (p *Product) GetID() int {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetBrandID() brands.ID {
	return p.Brand
}

func (p *Product) GetCategory() string {
	return p.Category
}

func (p *Product) GetPrice() prices.Price {
	return p.Price
}

func (p *Product) GetPriceAmount() int {
	return int(p.Price.Amount())
}

func (p *Product) GetSize() sizes.Size {
	return p.Size
}
