package categoriess

import (
	"fmt"
)

type Category string

const (
	Tops        Category = "TOPS"
	Outerwear   Category = "OUTERWEAR"
	Pants       Category = "PANTS"
	Sneakers    Category = "SNEAKERS"
	Bags        Category = "BAGS"
	Hats        Category = "HATS"
	Socks       Category = "SOCKS"
	Accessories Category = "ACCESSORIES"
)

var (
	CATEGORIES = []Category{
		Tops,
		Outerwear,
		Pants,
		Sneakers,
		Bags,
		Hats,
		Socks,
		Accessories,
	}
)

var validCategories = map[Category]bool{
	Tops:        true,
	Outerwear:   true,
	Pants:       true,
	Sneakers:    true,
	Bags:        true,
	Hats:        true,
	Socks:       true,
	Accessories: true,
}

type Categories struct {
	values []Category
}

func NewCategories(categories []string) (*Categories, error) {
	values := make([]Category, len(categories))
	for i, category := range categories {
		if !validCategories[Category(category)] {
			return nil, fmt.Errorf("invalid category: %s", category)
		}
		values[i] = Category(category)
	}
	return &Categories{values: values}, nil
}

func (c *Categories) Values() []string {
	values := make([]string, len(c.values))
	for i, category := range c.values {
		values[i] = string(category)
	}
	return values
}

func (c *Categories) Contains(category Category) bool {
	for _, v := range c.values {
		if v == category {
			return true
		}
	}
	return false
}

func (c *Categories) Add(category Category) error {
	if !validCategories[category] {
		return fmt.Errorf("invalid category: %s", category)
	}

	if c.Contains(category) {
		return fmt.Errorf("category already exists: %s", category)
	}

	c.values = append(c.values, category)
	return nil
}

func (c *Categories) Remove(category Category) error {
	if !validCategories[category] {
		return fmt.Errorf("invalid category: %s", category)
	}

	categories := make([]Category, 0, len(c.values)-1)
	for _, v := range c.values {
		if v != category {
			categories = append(categories, v)
		}
	}
	c.values = categories
	return nil
}

func IsValidCategory(category string) bool {
	return validCategories[Category(category)]
}
