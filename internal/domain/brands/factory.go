package brands

import (
	"context"
	"dresser/internal/domain/categoriess"
	"errors"
)

// Errors
var (
	ErrEmptyBrandName = errors.New("brand name cannot be empty")
)

// Factory defines the port for brand creation
type Factory interface {
	CreateBrand(ctx context.Context, name string, categories []string) (*Brand, error)
}

// DefaultFactory implements the Factory interface
type DefaultFactory struct {
	repository Repository
}

// NewFactory creates a new instance of DefaultFactory
func NewFactory(repository Repository) DefaultFactory {
	return DefaultFactory{
		repository: repository,
	}
}

// CreateBrand implements Factory interface
func (f *DefaultFactory) CreateBrand(ctx context.Context, name string, categories []string) (*Brand, error) {
	if name == "" {
		return nil, ErrEmptyBrandName
	}

	cats, err := categoriess.NewCategories(categories)
	if err != nil {
		return nil, err
	}

	nextID := f.repository.NextID()
	brand, err := New(int(nextID), name, *cats)
	if err != nil {
		return nil, err
	}

	return brand, nil
}
