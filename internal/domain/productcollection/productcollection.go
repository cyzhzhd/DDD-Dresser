package productcollection

import (
	"dresser/internal/domain/brands"
	"dresser/internal/domain/categoriess"
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

func (pc *ProductCollection) FilterByCategory(category string) *ProductCollection {
	filteredProducts := make([]*products.Product, 0)
	for _, product := range pc.Products {
		if product.Category == category {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return NewProductCollection(filteredProducts)
}

func (pc *ProductCollection) FilterByBrand(brandID brands.ID) *ProductCollection {
	filteredProducts := make([]*products.Product, 0)
	for _, product := range pc.Products {
		if product.Brand == brandID {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return NewProductCollection(filteredProducts)
}

func (pc *ProductCollection) GetLowestPriceProduct() *products.Product {
	var (
		lowestPriceProduct *products.Product
		lowestPrice        = 100000000
	)

	for _, product := range pc.Products {
		// todo: 화폐 단위 고려
		if product.Price.Amount() < lowestPrice {
			lowestPrice = product.Price.Amount()
			lowestPriceProduct = product
		}
	}
	return lowestPriceProduct
}
func (pc *ProductCollection) GetHighestPriceProduct() *products.Product {
	var (
		highestPriceProduct *products.Product
		highestPrice        = 0
	)

	for _, product := range pc.Products {
		// todo: 화폐 단위 고려
		if product.Price.Amount() > highestPrice {
			highestPrice = product.Price.Amount()
			highestPriceProduct = product
		}
	}
	return highestPriceProduct
}

func (pc *ProductCollection) GetLowestProductsByCategories() *ProductCollection {
	cs := categoriess.CATEGORIES

	var lowestProducts []*products.Product

	for _, c := range cs {
		p := pc.FilterByCategory(string(c)).GetLowestPriceProduct()
		lowestProducts = append(lowestProducts, p)
	}

	return NewProductCollection(lowestProducts)
}
