package productcollection

import (
	"context"
	"dresser/internal/domain/brands"
	"dresser/internal/domain/categoriess"
	"dresser/internal/domain/productcollection"
	"dresser/internal/domain/products"
)

type ICategoriesQueryApplicationService interface {
	GetLowestProducts(ctx context.Context) (*LowestProductsDTO, error)
}

type CategoriesQueryService struct {
	brandRepository   brands.Repository
	productRepository products.Repository
}

func NewCategoriesQueryService(brandRepository brands.Repository, productRepository products.Repository) *CategoriesQueryService {
	return &CategoriesQueryService{
		brandRepository:   brandRepository,
		productRepository: productRepository,
	}
}

// ProductDTO represents a product with its brand information
type ProductDTO struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Category  string `json:"category"`
	BrandID   int    `json:"brandId"`
	BrandName string `json:"brandName"`
}

// LowestProductsDTO represents the collection of lowest priced products per category
type LowestProductsDTO struct {
	Products   []ProductDTO `json:"products"`
	TotalPrice int          `json:"totalPrice"`
}

func (s *CategoriesQueryService) GetLowestProducts(ctx context.Context) (*LowestProductsDTO, error) {
	ps, err := s.productRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	cs := categoriess.CATEGORIES

	var lowestProducts []ProductDTO
	totalPrice := 0

	pc := productcollection.NewProductCollection(ps)

	for _, c := range cs {
		p := pc.FilterByCategory(string(c)).GetLowestPriceProduct()
		b, err := s.brandRepository.FindByID(ctx, p.GetBrandID())
		if err != nil {
			return nil, err
		}
		lowestProducts = append(lowestProducts, ProductDTO{
			ID:        p.GetID(),
			Name:      p.GetName(),
			Price:     p.GetPriceAmount(),
			Category:  string(c),
			BrandID:   b.GetID(),
			BrandName: b.GetName(),
		})
		totalPrice += p.GetPriceAmount()
	}

	// Convert to DTO
	return &LowestProductsDTO{
		Products:   lowestProducts,
		TotalPrice: totalPrice,
	}, nil
}
