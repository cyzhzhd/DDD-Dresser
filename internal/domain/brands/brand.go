package brands

type ID int

type Brand struct {
	ID   ID
	Name string
}

func New(id int, name string) (*Brand, error) {
	return &Brand{
		ID:   ID(id),
		Name: name,
	}, nil
}

func (b *Brand) GetID() ID {
	return b.ID
}

func (b *Brand) GetName() string {
	return b.Name
}
