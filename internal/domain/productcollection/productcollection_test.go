package productcollection

import (
	"dresser/internal/domain/brands"
	"dresser/internal/domain/categoriess"
	"dresser/internal/domain/prices"
	"dresser/internal/domain/products"
	"dresser/internal/domain/sizes"
	"encoding/json"
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
	mockData    *MockData
	productList []*products.Product
	testSuitePC *ProductCollection
)

// TestMain은 모든 테스트 실행 전에 한 번 호출되는 함수입니다.
func TestMain(m *testing.M) {
	// 테스트 전 준비: 모든 테스트에서 공유할 목 데이터 및 객체 설정
	setup()

	// 테스트 실행
	exitCode := m.Run()

	// 테스트 후 정리 (필요 시)
	// teardown()

	// 테스트 종료
	os.Exit(exitCode)
}

// 모든 테스트 실행 전에 필요한 데이터 로드 및 초기화
func setup() {
	var err error

	// 목 데이터 로드
	mockData, err = loadMockData()
	if err != nil {
		// 테스트 셋업 실패 시 출력하고 종료
		panic("Failed to load mock data: " + err.Error())
	}

	// 제품 목록 생성
	productList, err = createProducts(mockData)
	if err != nil {
		panic("Failed to create products: " + err.Error())
	}

	// ProductCollection 생성
	testSuitePC = NewProductCollection(productList)
}

func loadMockData() (*MockData, error) {
	data, err := os.ReadFile("../../../tools/mock_data.json")
	if err != nil {
		return nil, err
	}

	var mockData MockData
	if err := json.Unmarshal(data, &mockData); err != nil {
		return nil, err
	}

	return &mockData, nil
}

func createProducts(mockData *MockData) ([]*products.Product, error) {
	productList := make([]*products.Product, 0, len(mockData.Products))

	for i, p := range mockData.Products {
		price, err := prices.NewPrice(p.Amount, prices.Currency(p.Currency))
		if err != nil {
			return nil, err
		}

		size, err := sizes.NewSize(p.Size, categoriess.Category(p.Category))
		if err != nil {
			return nil, err
		}

		product := products.New(i+1, p.Name, brands.ID(p.BrandID), p.Category, *price, *size)
		productList = append(productList, product)
	}

	return productList, nil
}

func TestNewProductCollection(t *testing.T) {
	pc := NewProductCollection(productList)
	if pc == nil {
		t.Fatal("Expected non-nil ProductCollection")
	}

	if len(pc.Products) != len(mockData.Products) {
		t.Errorf("Expected %d products, got %d", len(mockData.Products), len(pc.Products))
	}
}

func TestFilterByCategory(t *testing.T) {
	// Count products with TOPS category in mock data
	topsCount := 0
	for _, p := range mockData.Products {
		if p.Category == string(categoriess.Tops) {
			topsCount++
		}
	}

	// Test filtering by TOPS category
	topsPC := testSuitePC.FilterByCategory(string(categoriess.Tops))
	if len(topsPC.Products) != topsCount {
		t.Errorf("Expected %d TOPS products, got %d", topsCount, len(topsPC.Products))
	}

	for _, p := range topsPC.Products {
		if p.Category != string(categoriess.Tops) {
			t.Errorf("Expected TOPS category, got %s", p.Category)
		}
	}
}

func TestFilterByBrand(t *testing.T) {
	// Count products with Brand ID 1 in mock data
	brandOneCount := 0
	for _, p := range mockData.Products {
		if p.BrandID == 1 {
			brandOneCount++
		}
	}

	// Test filtering by Brand ID 1
	brandOnePC := testSuitePC.FilterByBrand(brands.ID(1))
	if len(brandOnePC.Products) != brandOneCount {
		t.Errorf("Expected %d Brand 1 products, got %d", brandOneCount, len(brandOnePC.Products))
	}

	for _, p := range brandOnePC.Products {
		if p.Brand != brands.ID(1) {
			t.Errorf("Expected Brand ID 1, got %d", p.Brand)
		}
	}
}

func TestGetLowestPriceProduct(t *testing.T) {
	// Find the lowest price product in mock data
	lowestPrice := 100000000
	var lowestPriceProductName string
	for _, p := range mockData.Products {
		if p.Amount < lowestPrice {
			lowestPrice = p.Amount
			lowestPriceProductName = p.Name
		}
	}

	lowestPriceProduct, err := testSuitePC.GetLowestPriceProduct()
	if err != nil {
		t.Fatal("Expected non-nil lowest price product")
	}

	if lowestPriceProduct.GetName() != lowestPriceProductName {
		t.Errorf("Expected lowest price product %s, got %s", lowestPriceProductName, lowestPriceProduct.GetName())
	}

	if lowestPriceProduct.Price.Amount() != lowestPrice {
		t.Errorf("Expected lowest price %d, got %d", lowestPrice, lowestPriceProduct.Price.Amount())
	}
}

func TestGetHighestPriceProduct(t *testing.T) {
	// Find the highest price product in mock data
	highestPrice := 0
	var highestPriceProductName string
	for _, p := range mockData.Products {
		if p.Amount > highestPrice {
			highestPrice = p.Amount
			highestPriceProductName = p.Name
		}
	}

	highestPriceProduct, err := testSuitePC.GetHighestPriceProduct()
	if err != nil {
		t.Fatal("Expected non-nil highest price product")
	}

	if highestPriceProduct.GetName() != highestPriceProductName {
		t.Errorf("Expected highest price product %s, got %s", highestPriceProductName, highestPriceProduct.GetName())
	}

	if highestPriceProduct.Price.Amount() != highestPrice {
		t.Errorf("Expected highest price %d, got %d", highestPrice, highestPriceProduct.Price.Amount())
	}
}

func TestGetLowestProductsByCategories(t *testing.T) {
	lowestByCategory := testSuitePC.GetLowestProductsByCategories()
	if lowestByCategory == nil {
		t.Fatal("Expected non-nil result from GetLowestProductsByCategories")
	}

	assert.Equal(t, len(lowestByCategory.Products), 8)
	for _, p := range lowestByCategory.Products {
		switch p.GetCategory() {
		case string(categoriess.Tops):
			assert.Equal(t, p.GetPriceAmount(), 10000)
		case string(categoriess.Outerwear):
			assert.Equal(t, p.GetPriceAmount(), 5000)
		case string(categoriess.Pants):
			assert.Equal(t, p.GetPriceAmount(), 3000)
		case string(categoriess.Sneakers):
			assert.Equal(t, p.GetPriceAmount(), 9000)
		case string(categoriess.Bags):
			assert.Equal(t, p.GetPriceAmount(), 2000)
		case string(categoriess.Hats):
			assert.Equal(t, p.GetPriceAmount(), 1500)
		case string(categoriess.Socks):
			assert.Equal(t, p.GetPriceAmount(), 1700)
		case string(categoriess.Accessories):
			assert.Equal(t, p.GetPriceAmount(), 1900)
		default:
			assert.Equal(t, p, "")
		}
	}
}
