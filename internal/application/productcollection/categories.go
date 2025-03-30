package productcollection

import (
	"context"
	"dresser/internal/domain/brands"
	"dresser/internal/domain/categoriess"
	"dresser/internal/domain/productcollection"
	"dresser/internal/domain/products"
	"fmt"
)

type ICategoriesQueryApplicationService interface {
	GetLowestProducts(ctx context.Context) (*LowestProductsDTO, error)
	GetLowestProductsByBrand(ctx context.Context) (*LowestProductsByBrand, error)
	GetLowestAndHighestProductsByCategories(ctx context.Context, category string) (*LoHiProductsByCategory, error)
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

type ProductDTO struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Category  string `json:"category"`
	BrandID   int    `json:"brandId"`
	BrandName string `json:"brandName"`
}

type LowestProductsDTO struct {
	Products   []ProductDTO `json:"products"`
	TotalPrice int          `json:"totalPrice"`
}

// GetLowestProducts 메서드는 모든 카테고리에서 가장 저렴한 제품을 조회합니다.
func (s *CategoriesQueryService) GetLowestProducts(ctx context.Context) (*LowestProductsDTO, error) {
	ps, err := s.productRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	pc := productcollection.NewProductCollection(ps)
	// 카테고리 별로 가장 저렴한 제품을 조회합니다.
	lps := pc.GetLowestProductsByCategories()

	var lowestProducts []ProductDTO
	totalPrice := 0

	// DTO로 변환하며 전체 가격을 추가합니다.
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

// GetLowestProductsByBrand 메서드는 모든 카테고리를 가장 저렴하게 살 수 있는 브랜드를 조회합니다.
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

	// 브랜드의 카테고리별 가장 저렴한 제품을 고릅니다.
	for _, b := range bs {
		brandProducts := pc.FilterByBrand(b.GetID())

		// 카테고리별 가장 저렴한 제품들의 목록을 가져옵니다.
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

		candidates = append(candidates, &LowestProductsDTO{
			Products:   dtos,
			TotalPrice: totalPrice,
		})
	}

	// 전체 금액이 가장 저렴한 브랜드를 찾습니다.
	minTotalPrice := candidates[0].TotalPrice
	minTotalPriceCandidate := candidates[0]

	for _, c := range candidates {
		if c.TotalPrice < minTotalPrice {
			minTotalPrice = c.TotalPrice
			minTotalPriceCandidate = c
		}
	}
	fmt.Println("minTotalPrice", minTotalPrice)
	fmt.Println("minTotalPriceCandidate", minTotalPriceCandidate)

	return &LowestProductsByBrand{
		Lowest: LowestProductsBrandDTO{
			Brand:      minTotalPriceCandidate.Products[0].BrandName,
			Category:   minTotalPriceCandidate.Products,
			TotalPrice: minTotalPriceCandidate.TotalPrice,
		},
	}, nil
}

type LoHiProductsByCategory struct {
	Category string
	Lowest   ProductDTO
	Highest  ProductDTO
}

// GetLowestAndHighestProductsByCategories 메서드는 카테고리별 가장 저렴한 제품과 가장 비싼 제품을 조회합니다.
func (s *CategoriesQueryService) GetLowestAndHighestProductsByCategories(ctx context.Context, category string) (*LoHiProductsByCategory, error) {
	if !categoriess.IsValidCategory(category) {
		return nil, fmt.Errorf("invalid category: %s", category)
	}

	ps, err := s.productRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	pc := productcollection.NewProductCollection(ps).FilterByCategory(category)

	lowest, err := pc.GetLowestPriceProduct()
	if err != nil {
		return nil, err
	}
	highest, err := pc.GetHighestPriceProduct()
	if err != nil {
		return nil, err
	}

	lowestBrand, err := s.brandRepository.FindByID(ctx, lowest.GetBrandID())
	if err != nil {
		return nil, err
	}
	highestBrand, err := s.brandRepository.FindByID(ctx, highest.GetBrandID())
	if err != nil {
		return nil, err
	}
	lowestDTO := ProductDTO{
		ID:        lowest.GetID(),
		Name:      lowest.GetName(),
		Price:     lowest.GetPriceAmount(),
		Category:  string(category),
		BrandID:   int(lowestBrand.GetID()),
		BrandName: lowestBrand.GetName(),
	}

	highestDTO := ProductDTO{
		ID:        highest.GetID(),
		Name:      highest.GetName(),
		Price:     highest.GetPriceAmount(),
		Category:  string(category),
		BrandID:   int(highestBrand.GetID()),
		BrandName: highestBrand.GetName(),
	}

	return &LoHiProductsByCategory{
		Category: string(category),
		Lowest:   lowestDTO,
		Highest:  highestDTO,
	}, nil
}
