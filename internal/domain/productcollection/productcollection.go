package productcollection

import (
	"dresser/internal/domain/brands"
	"dresser/internal/domain/products"
)

type ProductCollection struct {
	Products []*products.Product
}

func NewProductCollection(products []*products.Product) *ProductCollection {
	return &ProductCollection{
		Products: products,
	}
}

func (c *ProductCollection) FilterByCategory(category string) *ProductCollection {
	filteredProducts := make([]*products.Product, 0)
	for _, product := range c.Products {
		if product.Category == category {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return NewProductCollection(filteredProducts)
}

func (c *ProductCollection) FilterByBrand(brandID brands.ID) *ProductCollection {
	filteredProducts := make([]*products.Product, 0)
	for _, product := range c.Products {
		if product.Brand == brandID {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return NewProductCollection(filteredProducts)
}

func (c *ProductCollection) GetLowestPriceProduct() *products.Product {
	var (
		lowestPriceProduct *products.Product
		lowestPrice        = 100000000
	)

	for _, product := range c.Products {
		// todo: 화폐 단위 고려
		if product.Price.Amount() < lowestPrice {
			lowestPrice = product.Price.Amount()
			lowestPriceProduct = product
		}
	}
	return lowestPriceProduct
}
