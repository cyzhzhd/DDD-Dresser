package sizes

import (
	"dresser/internal/domain/categoriess"
	"fmt"
)

// SizeType represents the type of size measurement system
type SizeType string

const (
	// Clothing sizes (S, M, L, XL, etc.)
	LetterSize SizeType = "LETTER"
	// Numeric sizes (36, 37, 38, etc.)
	NumericSize SizeType = "NUMERIC"
	// Dimensional sizes (28x30, 30x32, etc.)
	DimensionalSize SizeType = "DIMENSIONAL"
	// Free size
	FreeSize SizeType = "FREE"
)

// CategorySizeMap defines valid size types for each category
var CategorySizeMap = map[categoriess.Category]SizeType{
	categoriess.Tops:        LetterSize,
	categoriess.Outerwear:   LetterSize,
	categoriess.Pants:       DimensionalSize,
	categoriess.Sneakers:    NumericSize,
	categoriess.Bags:        FreeSize,
	categoriess.Hats:        LetterSize,
	categoriess.Socks:       LetterSize,
	categoriess.Accessories: FreeSize,
}

// SizeSpecification defines the interface for size validation
type SizeSpecification interface {
	IsSatisfiedBy(value string) bool
	Error(value string) error
}

// LetterSizeSpec implements letter size validation
type LetterSizeSpec struct {
	validSizes map[string]bool
}

func NewLetterSizeSpec() *LetterSizeSpec {
	return &LetterSizeSpec{
		validSizes: map[string]bool{
			"XS": true, "S": true, "M": true, "L": true, "XL": true, "XXL": true,
		},
	}
}

func (s *LetterSizeSpec) IsSatisfiedBy(value string) bool {
	return s.validSizes[value]
}

func (s *LetterSizeSpec) Error(value string) error {
	return fmt.Errorf("invalid letter size: %s", value)
}

// NumericSizeSpec implements numeric size validation
type NumericSizeSpec struct {
	minSize int
	maxSize int
}

func NewNumericSizeSpec() *NumericSizeSpec {
	return &NumericSizeSpec{
		minSize: 220,
		maxSize: 300,
	}
}

func (s *NumericSizeSpec) IsSatisfiedBy(value string) bool {
	var size int
	if _, err := fmt.Sscanf(value, "%d", &size); err != nil {
		return false
	}
	return size >= s.minSize && size <= s.maxSize
}

func (s *NumericSizeSpec) Error(value string) error {
	return fmt.Errorf("invalid numeric size: %s (must be between %d and %d)", value, s.minSize, s.maxSize)
}

// DimensionalSizeSpec implements dimensional size validation
type DimensionalSizeSpec struct {
	minWaist  int
	maxWaist  int
	minLength int
	maxLength int
}

func NewDimensionalSizeSpec() *DimensionalSizeSpec {
	return &DimensionalSizeSpec{
		minWaist:  26,
		maxWaist:  40,
		minLength: 28,
		maxLength: 36,
	}
}

func (s *DimensionalSizeSpec) IsSatisfiedBy(value string) bool {
	var waist, length int
	if _, err := fmt.Sscanf(value, "%dx%d", &waist, &length); err != nil {
		return false
	}
	return waist >= s.minWaist && waist <= s.maxWaist &&
		length >= s.minLength && length <= s.maxLength
}

func (s *DimensionalSizeSpec) Error(value string) error {
	return fmt.Errorf("invalid dimensional size: %s (waist: %d-%d, length: %d-%d)",
		value, s.minWaist, s.maxWaist, s.minLength, s.maxLength)
}

// FreeSizeSpec implements free size validation
type FreeSizeSpec struct{}

func NewFreeSizeSpec() *FreeSizeSpec {
	return &FreeSizeSpec{}
}

func (s *FreeSizeSpec) IsSatisfiedBy(value string) bool {
	return value == "FREE"
}

func (s *FreeSizeSpec) Error(value string) error {
	return fmt.Errorf("invalid free size value: %s (must be 'FREE')", value)
}

// Size represents a value object for product sizes
type Size struct {
	value    string
	sizeType SizeType
}

var sizeSpecs = map[SizeType]SizeSpecification{
	LetterSize:      NewLetterSizeSpec(),
	NumericSize:     NewNumericSizeSpec(),
	DimensionalSize: NewDimensionalSizeSpec(),
	FreeSize:        NewFreeSizeSpec(),
}

// NewSize creates a new Size value object with validation
func NewSize(value string, category categoriess.Category) (*Size, error) {
	sizeType, exists := CategorySizeMap[category]
	if !exists {
		return nil, fmt.Errorf("invalid category for size: %s", category)
	}

	spec, exists := sizeSpecs[sizeType]
	if !exists {
		return nil, fmt.Errorf("no specification found for size type: %s", sizeType)
	}

	if !spec.IsSatisfiedBy(value) {
		return nil, spec.Error(value)
	}

	return &Size{
		value:    value,
		sizeType: sizeType,
	}, nil
}

// Value returns the size value
func (s *Size) Value() string {
	return s.value
}

// Type returns the size type
func (s *Size) Type() SizeType {
	return s.sizeType
}
