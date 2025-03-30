package brands

import "dresser/internal/domain/categoriess"

type ID int

type Brand struct {
	ID         ID
	Name       string
	Categories categoriess.Categories
}

func New(id int, name string, categories categoriess.Categories) (*Brand, error) {
	return &Brand{
		ID:         ID(id),
		Name:       name,
		Categories: categories,
	}, nil
}

func (b *Brand) AddCategory(category categoriess.Category) error {
	return b.Categories.Add(category)
}

func (b *Brand) RemoveCategory(category categoriess.Category) error {
	return b.Categories.Remove(category)
}

func (b *Brand) GetID() int {
	return int(b.ID)
}

func (b *Brand) GetName() string {
	return b.Name
}

func (b *Brand) GetCategories() []string {
	return b.Categories.Values()
}
