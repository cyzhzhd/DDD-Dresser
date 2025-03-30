package productcollection

import (
	"context"
	"dresser/internal/domain/brands"
	"dresser/internal/domain/categoriess"
	"dresser/internal/domain/prices"
	"dresser/internal/domain/products"
	"dresser/internal/domain/sizes"
	"dresser/internal/infrastructure/db"
	pbrands "dresser/internal/infrastructure/persistence/brands"
	pproducts "dresser/internal/infrastructure/persistence/products"
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockProduct struct {
	Name     string `json:"name"`
	BrandID  int    `json:"brand_id"`
	Category string `json:"category"`
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
	Size     string `json:"size"`
}

type MockData struct {
	Brands   []struct{ Name string } `json:"brands"`
	Products []MockProduct           `json:"products"`
}

// 테스트 스위트에서 공유할 변수들
var (
	mockData          *MockData
	mockProductList   []*products.Product
	mockBrandList     []*brands.Brand
	mockCtx           context.Context
	categoriesService *CategoriesQueryService
	database          db.DB
	productRepo       products.Repository
	brandRepo         brands.Repository
)

// TestMain은 모든 테스트 실행 전에 한 번 호출되는 함수입니다.
func TestMain(m *testing.M) {
	// 테스트 전 준비: 모든 테스트에서 공유할 목 데이터 및 객체 설정
	setup()

	// 테스트 실행
	exitCode := m.Run()

	// 테스트 종료
	if database != nil {
		database.Close()
	}
	os.Exit(exitCode)
}

// 모든 테스트 실행 전에 필요한 데이터 로드 및 초기화
func setup() {
	var err error

	// 컨텍스트 초기화
	mockCtx = context.Background()

	// 목 데이터 로드
	mockData, err = loadMockData()
	if err != nil {
		panic("Failed to load mock data: " + err.Error())
	}

	// 제품 및 브랜드 목록 생성
	mockProductList, mockBrandList, err = createEntities(mockData)
	if err != nil {
		panic("Failed to create entities: " + err.Error())
	}

	// 메모리 데이터베이스 초기화
	database, err = db.New(db.Config{
		Type: db.Memory,
	})
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 리포지토리 초기화
	brandRepo = pbrands.NewBrandRepository(database)
	productRepo = pproducts.NewProductRepository(database)

	// 목 데이터를 데이터베이스에 저장
	for _, brand := range mockBrandList {
		_, err := brandRepo.Create(mockCtx, brand)
		if err != nil {
			log.Fatalf("Failed to save brand to repository: %v", err)
		}
	}

	for _, product := range mockProductList {
		_, err := productRepo.Create(mockCtx, product)
		if err != nil {
			log.Fatalf("Failed to save product to repository: %v", err)
		}
	}

	// 서비스 초기화
	categoriesService = NewCategoriesQueryService(brandRepo, productRepo)
}

func loadMockData() (*MockData, error) {
	data, err := os.ReadFile("../../../tools/mock_data_small.json")
	if err != nil {
		return nil, err
	}

	var mockData MockData
	if err := json.Unmarshal(data, &mockData); err != nil {
		return nil, err
	}

	return &mockData, nil
}

func createEntities(mockData *MockData) ([]*products.Product, []*brands.Brand, error) {
	productList := make([]*products.Product, 0, len(mockData.Products))
	brandList := make([]*brands.Brand, 0, len(mockData.Brands))

	// 브랜드 생성
	for i, b := range mockData.Brands {
		brand, err := brands.New(i+1, b.Name)
		if err != nil {
			return nil, nil, err
		}
		brandList = append(brandList, brand)
	}

	// 제품 생성
	for i, p := range mockData.Products {
		price, err := prices.NewPrice(p.Amount, prices.Currency(p.Currency))
		if err != nil {
			return nil, nil, err
		}

		size, err := sizes.NewSize(p.Size, categoriess.Category(p.Category))
		if err != nil {
			return nil, nil, err
		}

		product := products.New(i+1, p.Name, brands.ID(p.BrandID), p.Category, *price, *size)
		productList = append(productList, product)
	}

	return productList, brandList, nil
}

func TestGetLowestProducts(t *testing.T) {
	result, err := categoriesService.GetLowestProducts(mockCtx)
	assert.NoError(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, 3, len(result.Products))
	assert.Equal(t, "C", result.Products[0].BrandName)
	assert.Equal(t, 18100, result.TotalPrice)
}

func TestGetLowestProductsByBrand(t *testing.T) {
	result, err := categoriesService.GetLowestProductsByBrand(mockCtx)
	assert.NoError(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, "D", result.Lowest.Brand)
	assert.Equal(t, 18200, result.Lowest.TotalPrice)
}

func TestGetLowestAndHighestProductsByCategories(t *testing.T) {
	// TOPS 카테고리에 대한 테스트
	category := string(categoriess.Tops)
	result, err := categoriesService.GetLowestAndHighestProductsByCategories(mockCtx, category)
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// 카테고리 검증
	assert.Equal(t, category, result.Category)

	assert.Equal(t, "C Tops", result.Lowest.Name)
	assert.Equal(t, 10000, result.Lowest.Price)
	assert.Equal(t, "A Tops", result.Highest.Name)
	assert.Equal(t, 11200, result.Highest.Price)
}

func TestGetLowestAndHighestProductsByInvalidCategory(t *testing.T) {
	// 잘못된 카테고리로 테스트
	invalidCategory := "INVALID_CATEGORY"
	result, err := categoriesService.GetLowestAndHighestProductsByCategories(mockCtx, invalidCategory)
	assert.Error(t, err)
	assert.Nil(t, result)
}
