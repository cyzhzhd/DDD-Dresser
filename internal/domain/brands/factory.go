package brands

import (
	"context"
	"errors"
)

// Errors
var (
	ErrEmptyBrandName = errors.New("brand name cannot be empty")
)

// Factory defines the port for brand creation
type Factory interface {
	CreateBrand(ctx context.Context, name string) (*Brand, error)
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
func (f *DefaultFactory) CreateBrand(ctx context.Context, name string) (*Brand, error) {
	if name == "" {
		return nil, ErrEmptyBrandName
	}

	nextID := f.repository.NextID()
	brand, err := New(int(nextID), name)
	if err != nil {
		return nil, err
	}

	return brand, nil
}
