package productcollection

import (
	"dresser/internal/domain/brands"
	"dresser/internal/domain/categoriess"
	"dresser/internal/domain/products"
	"errors"
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

func (pc *ProductCollection) GetLowestPriceProduct() (*products.Product, error) {
	var (
		lowestPriceProduct *products.Product
		lowestPrice        = 100000000
	)
	if len(pc.Products) == 0 {
		return nil, errors.New("no products")
	}

	for _, product := range pc.Products {
		// todo: 화폐 단위 고려
		if product.Price.Amount() < lowestPrice {
			lowestPrice = product.Price.Amount()
			lowestPriceProduct = product
		}
	}
	return lowestPriceProduct, nil
}
func (pc *ProductCollection) GetHighestPriceProduct() (*products.Product, error) {
	var (
		highestPriceProduct *products.Product
		highestPrice        = 0
	)
	if len(pc.Products) == 0 {
		return nil, errors.New("no products")
	}

	for _, product := range pc.Products {
		// todo: 화폐 단위 고려
		if product.Price.Amount() > highestPrice {
			highestPrice = product.Price.Amount()
			highestPriceProduct = product
		}
	}
	return highestPriceProduct, nil
}

func (pc *ProductCollection) GetLowestProductsByCategories() *ProductCollection {
	cs := categoriess.CATEGORIES

	var lowestProducts []*products.Product

	for _, c := range cs {
		p, err := pc.FilterByCategory(string(c)).GetLowestPriceProduct()
		if err != nil {
			continue
		}
		lowestProducts = append(lowestProducts, p)
	}

	return NewProductCollection(lowestProducts)
}
