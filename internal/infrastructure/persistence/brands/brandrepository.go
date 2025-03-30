package brands

import (
	"context"
	"dresser/internal/domain/brands"
	"dresser/internal/domain/categoriess"
	"dresser/internal/infrastructure/db"
	"fmt"
	"sync"
)

// BrandRepository implements application.Repository interface
type BrandRepository struct {
	sync.RWMutex
	brands map[brands.ID]*brands.Brand
	lastID brands.ID
	db     db.DB
}

// NewBrandRepository creates a new instance of BrandRepository with a database
func NewBrandRepository(database db.DB) brands.Repository {
	return &BrandRepository{
		brands: make(map[brands.ID]*brands.Brand),
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
	res, err := r.db.Create(ctx, "brands", brand.GetID(), brand)
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
	// Use database if available
	if r.db != nil {
		return r.db.Update(ctx, "brands", int(brand.ID), brand)
	}

	// Fallback to in-memory implementation
	r.Lock()
	defer r.Unlock()

	if _, exists := r.brands[brand.ID]; !exists {
		return fmt.Errorf("brand with ID %d not found", brand.ID)
	}

	brandCopy := *brand
	r.brands[brandCopy.ID] = &brandCopy

	return nil
}

// Delete implements brands.Repository
func (r *BrandRepository) Delete(ctx context.Context, id brands.ID) error {
	// Use database if available
	if r.db != nil {
		return r.db.Delete(ctx, "brands", int(id))
	}

	// Fallback to in-memory implementation
	r.Lock()
	defer r.Unlock()

	if _, exists := r.brands[id]; !exists {
		return fmt.Errorf("brand with ID %d not found", id)
	}

	delete(r.brands, id)
	return nil
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
	// Use database if available
	if r.db != nil {
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

	// Fallback to in-memory implementation
	r.RLock()
	defer r.RUnlock()

	var result []*brands.Brand
	for _, brand := range r.brands {
		if brand.Categories.Contains(categoriess.Category(category)) {
			result = append(result, brand)
		}
	}

	return result, nil
}

// Exists implements brands.Repository
func (r *BrandRepository) Exists(ctx context.Context, id brands.ID) (bool, error) {
	// Use database if available
	if r.db != nil {
		// BrandExists calls FindByID internally, so we can just use that
		_, err := r.db.FindByID(ctx, "brands", int(id))
		if err != nil {
			return false, nil // Not found
		}
		return true, nil
	}

	// Fallback to in-memory implementation
	r.RLock()
	defer r.RUnlock()

	_, exists := r.brands[id]
	return exists, nil
}
