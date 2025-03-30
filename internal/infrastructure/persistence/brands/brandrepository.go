package brands

import (
	"context"
	"dresser/internal/domain/brands"
	"dresser/internal/infrastructure/db"
	"fmt"
	"sync"
)

// BrandRepository implements application.Repository interface
type BrandRepository struct {
	sync.RWMutex
	lastID brands.ID
	db     db.DB
}

// NewBrandRepository creates a new instance of BrandRepository with a database
func NewBrandRepository(database db.DB) brands.Repository {
	return &BrandRepository{
		lastID: 0,
		db:     database,
	}
}

// NextID implements brands.Repository
func (r *BrandRepository) NextID() brands.ID {
	return brands.ID(r.db.NextID("brands"))
}

// Create implements brands.Repository
func (r *BrandRepository) Create(ctx context.Context, brand *brands.Brand) (*brands.Brand, error) {
	res, err := r.db.Create(ctx, "brands", int(brand.GetID()), brand)
	if err != nil {
		return nil, err
	}

	return res.(*brands.Brand), nil
}

// FindByID implements brands.Repository
func (r *BrandRepository) FindByID(ctx context.Context, id brands.ID) (*brands.Brand, error) {
	entity, err := r.db.FindByID(ctx, "brands", int(id))
	if err != nil {
		return nil, err
	}

	if entity == nil {
		return nil, fmt.Errorf("brand with ID %d not found", id)
	}

	if domainBrand, ok := entity.(*brands.Brand); ok {
		return domainBrand, nil
	}

	return nil, fmt.Errorf("fail to find brand with ID %d", id)
}

// Update implements brands.Repository
func (r *BrandRepository) Update(ctx context.Context, brand *brands.Brand) error {
	return r.db.Update(ctx, "brands", int(brand.ID), brand)
}

// Delete implements brands.Repository
func (r *BrandRepository) Delete(ctx context.Context, id brands.ID) error {
	return r.db.Delete(ctx, "brands", int(id))
}

// FindAll implements brands.Repository
func (r *BrandRepository) FindAll(ctx context.Context) ([]*brands.Brand, error) {
	entities, err := r.db.FindAll(ctx, "brands")
	if err != nil {
		return nil, err
	}
	fmt.Println("entities", entities)

	result := make([]*brands.Brand, 0, len(entities))
	for _, entity := range entities {
		if domainBrand, ok := entity.(*brands.Brand); ok {
			result = append(result, domainBrand)
		}
	}

	return result, nil
}

// FindByCategory implements brands.Repository
func (r *BrandRepository) FindByCategory(ctx context.Context, category string) ([]*brands.Brand, error) {
	filter := map[string]interface{}{
		"category": category,
	}

	entities, err := r.db.FindBy(ctx, "brands", filter)
	if err != nil {
		return nil, err
	}

	result := make([]*brands.Brand, 0, len(entities))
	for _, entity := range entities {
		if domainBrand, ok := entity.(*brands.Brand); ok {
			result = append(result, domainBrand)
		}
	}

	return result, nil
}

// Exists implements brands.Repository
func (r *BrandRepository) Exists(ctx context.Context, id brands.ID) (bool, error) {
	// BrandExists calls FindByID internally, so we can just use that
	_, err := r.db.FindByID(ctx, "brands", int(id))
	if err != nil {
		return false, nil // Not found
	}
	return true, nil
}
