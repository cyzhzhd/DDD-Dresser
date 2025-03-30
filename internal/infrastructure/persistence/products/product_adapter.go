package products

import (
	"context"
	appProducts "dresser/internal/application/products"
	"dresser/internal/domain/brands"
	"dresser/internal/domain/categoriess"
	"dresser/internal/domain/prices"
	dproducts "dresser/internal/domain/products"
	"dresser/internal/domain/sizes"
)

// ProductAdapter converts between domain and application product models
type ProductAdapter struct {
	registerService appProducts.IProductRegisterApplicationService
	updateService   appProducts.IProductUpdateApplicationService
	deleteService   appProducts.IProductDeleteApplicationService
	queryService    appProducts.IProductQueryApplicationService
}

// NewProductAdapter creates a new ProductAdapter
func NewProductAdapter(
	repository dproducts.Repository,
	factory dproducts.Factory,
) *ProductAdapter {
	return &ProductAdapter{
		registerService: appProducts.NewProductRegisterService(repository, factory),
		updateService:   appProducts.NewProductUpdateService(repository, factory),
		deleteService:   appProducts.NewProductDeleteService(repository),
		queryService:    appProducts.NewProductQueryService(repository),
	}
}

// ToDomain converts application product to domain product
func (a *ProductAdapter) ToDomain(appProduct *appProducts.ProductDTO) (*dproducts.Product, error) {
	// Create price
	price, err := prices.NewPrice(appProduct.Price.Amount, prices.Currency(appProduct.Price.Currency))
	if err != nil {
		return nil, err
	}

	// Create size
	size, err := sizes.NewSize(appProduct.Size.Value, categoriess.Category(appProduct.Category))
	if err != nil {
		return nil, err
	}

	return dproducts.New(
		appProduct.ID,
		appProduct.Name,
		brands.ID(appProduct.BrandID),
		appProduct.Category,
		*price,
		*size,
	), nil
}

// ToApplication converts domain product to application product
func (a *ProductAdapter) ToApplication(domainProduct *dproducts.Product) *appProducts.ProductDTO {
	price := domainProduct.GetPrice()
	size := domainProduct.GetSize()

	return &appProducts.ProductDTO{
		ID:       domainProduct.GetID(),
		Name:     domainProduct.GetName(),
		BrandID:  int(domainProduct.GetBrandID()),
		Category: domainProduct.GetCategory(),
		Price: appProducts.PriceDTO{
			Amount:       price.Amount(),
			Currency:     string(price.Currency()),
			DisplayValue: price.String(),
		},
		Size: appProducts.SizeDTO{
			Value: size.Value(),
			Type:  string(size.Type()),
		},
	}
}

// Register registers a new product
func (a *ProductAdapter) Register(ctx context.Context, cmd appProducts.RegisterProductCommand) (*appProducts.ProductDTO, error) {
	return a.registerService.Register(ctx, cmd)
}

// Update updates an existing product
func (a *ProductAdapter) Update(ctx context.Context, cmd appProducts.UpdateProductCommand) (*appProducts.ProductDTO, error) {
	return a.updateService.Update(ctx, cmd)
}

// Delete deletes a product by ID
func (a *ProductAdapter) Delete(ctx context.Context, id int) error {
	return a.deleteService.Delete(ctx, id)
}

// GetAll retrieves all products
func (a *ProductAdapter) GetAll(ctx context.Context) ([]*appProducts.ProductDTO, error) {
	return a.queryService.GetAll(ctx)
}

// Get retrieves a product by ID
func (a *ProductAdapter) Get(ctx context.Context, id int) (*appProducts.ProductDTO, error) {
	return a.queryService.Get(ctx, id)
}

// GetByBrand retrieves all products for a specific brand
func (a *ProductAdapter) GetByBrand(ctx context.Context, brandID int) ([]*appProducts.ProductDTO, error) {
	return a.queryService.GetByBrand(ctx, brandID)
}

// GetByCategory retrieves all products that have the specified category
func (a *ProductAdapter) GetByCategory(ctx context.Context, category string) ([]*appProducts.ProductDTO, error) {
	return a.queryService.GetByCategory(ctx, category)
}
