package products

import (
	"context"
	"database/sql"
	dproducts "dresser/internal/domain/products"
	"dresser/internal/infrastructure/db"
	"fmt"
	"sync"
)

// ProductRepository implements the domain product repository interface
type ProductRepository struct {
	sync.RWMutex
	db db.DB
}

// NewProductRepository creates a new instance of ProductRepository with a database
func NewProductRepository(database db.DB) dproducts.Repository {
	return &ProductRepository{
		db: database,
	}
}

// NextID implements dproducts.Repository
func (r *ProductRepository) NextID() int {
	return r.db.NextID("products")
}

// Create implements dproducts.Repository
func (r *ProductRepository) Create(ctx context.Context, product *dproducts.Product) (*dproducts.Product, error) {
	res, err := r.db.Create(ctx, "products", product.GetID(), product)
	if err != nil {
		return nil, err
	}

	return res.(*dproducts.Product), nil
}

// Update implements dproducts.Repository
func (r *ProductRepository) Update(ctx context.Context, product *dproducts.Product) error {
	return r.db.Update(ctx, "products", product.GetID(), product)
}

// Delete implements dproducts.Repository
func (r *ProductRepository) Delete(ctx context.Context, id int) error {
	return r.db.Delete(ctx, "products", id)
}

// FindByID implements dproducts.Repository
func (r *ProductRepository) FindByID(ctx context.Context, id int) (*dproducts.Product, error) {
	entity, err := r.db.FindByID(ctx, "products", id)
	if err != nil {
		return nil, err
	}

	if entity == nil {
		return nil, sql.ErrNoRows
	}

	if domainProduct, ok := entity.(*dproducts.Product); ok {
		return domainProduct, nil
	}

	return nil, fmt.Errorf("fail to find product with ID %d", id)
}

// FindAll implements dproducts.Repository
func (r *ProductRepository) FindAll(ctx context.Context) ([]*dproducts.Product, error) {
	entities, err := r.db.FindAll(ctx, "products")
	if err != nil {
		return nil, err
	}

	result := make([]*dproducts.Product, 0, len(entities))
	for _, entity := range entities {
		if domainProduct, ok := entity.(*dproducts.Product); ok {
			result = append(result, domainProduct)
		}
	}

	return result, nil
}

// FindByBrand implements dproducts.Repository
func (r *ProductRepository) FindByBrand(ctx context.Context, brandID int) ([]*dproducts.Product, error) {
	filter := map[string]interface{}{
		"brand": brandID,
	}

	entities, err := r.db.FindBy(ctx, "products", filter)
	if err != nil {
		return nil, err
	}

	result := make([]*dproducts.Product, 0, len(entities))
	for _, entity := range entities {
		if domainProduct, ok := entity.(*dproducts.Product); ok {
			result = append(result, domainProduct)
		}
	}

	return result, nil
}

// FindByCategory implements dproducts.Repository
func (r *ProductRepository) FindByCategory(ctx context.Context, category string) ([]*dproducts.Product, error) {
	filter := map[string]interface{}{
		"category": category,
	}

	entities, err := r.db.FindBy(ctx, "products", filter)
	if err != nil {
		return nil, err
	}

	result := make([]*dproducts.Product, 0, len(entities))
	for _, entity := range entities {
		if domainProduct, ok := entity.(*dproducts.Product); ok {
			result = append(result, domainProduct)
		}
	}

	return result, nil
}

// Exists implements dproducts.Repository
func (r *ProductRepository) Exists(ctx context.Context, id int) (bool, error) {
	_, err := r.db.FindByID(ctx, "products", id)
	if err != nil {
		return false, nil // Not found
	}
	return true, nil
}
