package productcollection

import (
	"context"
	"dresser/internal/domain/brands"
	"dresser/internal/domain/productcollection"
	"dresser/internal/domain/products"
)

type ICategoriesQueryApplicationService interface {
	GetLowestProducts(ctx context.Context) (*LowestProductsDTO, error)
	GetLowestProductsByBrand(ctx context.Context) (*LowestProductsByBrand, error)
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

	pc := productcollection.NewProductCollection(ps)
	lps := pc.GetLowestProductsByCategories()

	var lowestProducts []ProductDTO
	totalPrice := 0

	for _, lp := range lps.Products {
		b, err := s.brandRepository.FindByID(ctx, lp.GetBrandID())
		if err != nil {
			return nil, err
		}
		lowestProducts = append(lowestProducts, ProductDTO{
			ID:        lp.GetID(),
			Name:      lp.GetName(),
			Price:     lp.GetPriceAmount(),
			Category:  lp.GetCategory(),
			BrandID:   int(b.GetID()),
			BrandName: b.GetName(),
		})
		totalPrice += lp.GetPriceAmount()
	}

	// Convert to DTO
	return &LowestProductsDTO{
		Products:   lowestProducts,
		TotalPrice: totalPrice,
	}, nil
}

type LowestProductsByBrand struct {
	Lowest LowestProductsBrandDTO
}

type LowestProductsBrandDTO struct {
	Brand      string
	Category   []ProductDTO
	TotalPrice int
}

func (s *CategoriesQueryService) GetLowestProductsByBrand(ctx context.Context) (*LowestProductsByBrand, error) {
	ps, err := s.productRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	pc := productcollection.NewProductCollection(ps)

	bs, err := s.brandRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	candidates := make([]*LowestProductsDTO, 0)

	for _, b := range bs {
		brandProducts := pc.FilterByBrand(b.GetID())

		lps := brandProducts.GetLowestProductsByCategories()
		totalPrice := 0
		dtos := make([]ProductDTO, 0)

		for _, lp := range lps.Products {
			dtos = append(dtos, ProductDTO{
				ID:        lp.GetID(),
				Name:      lp.GetName(),
				Price:     lp.GetPriceAmount(),
				Category:  lp.GetCategory(),
				BrandID:   int(b.GetID()),
				BrandName: b.GetName(),
			})
			totalPrice += lp.GetPriceAmount()
		}

		// Convert to DTO
		candidates = append(candidates, &LowestProductsDTO{
			Products:   dtos,
			TotalPrice: totalPrice,
		})
	}

	minTotalPrice := candidates[0].TotalPrice
	minTotalPriceCandidate := candidates[0]

	for _, c := range candidates {
		if c.TotalPrice < minTotalPrice {
			minTotalPrice = c.TotalPrice
			minTotalPriceCandidate = c
		}
	}

	return &LowestProductsByBrand{
		Lowest: LowestProductsBrandDTO{
			Brand:      minTotalPriceCandidate.Products[0].BrandName,
			Category:   minTotalPriceCandidate.Products,
			TotalPrice: minTotalPriceCandidate.TotalPrice,
		},
	}, nil
}
