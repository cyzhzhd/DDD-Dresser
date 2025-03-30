package brands

import (
	"context"
	"errors"
)

var (
	ErrEmptyBrandName = errors.New("brand name cannot be empty")
)

type Factory interface {
	CreateBrand(ctx context.Context, name string) (*Brand, error)
}

type DefaultFactory struct {
	repository Repository
}

func NewFactory(repository Repository) DefaultFactory {
	return DefaultFactory{
		repository: repository,
	}
}

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
