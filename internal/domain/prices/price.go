package prices

import (
	"fmt"
	"math"
)

// Currency represents supported currencies
type Currency string

const (
	KRW Currency = "KRW"
	USD Currency = "USD"
	EUR Currency = "EUR"
)

var validCurrencies = map[Currency]bool{
	KRW: true,
	USD: true,
	EUR: true,
}

// Price represents a value object for monetary values
type Price struct {
	amount   int      // 금액 (최소 단위로 저장, e.g., KRW: 원, USD: cents)
	currency Currency // 통화
}

// NewPrice creates a new Price value object with validation
func NewPrice(amount int, currency Currency) (*Price, error) {
	if amount < 0 {
		return nil, fmt.Errorf("price amount cannot be negative: %d", amount)
	}

	if !validCurrencies[currency] {
		return nil, fmt.Errorf("invalid currency: %s", currency)
	}

	return &Price{
		amount:   amount,
		currency: currency,
	}, nil
}

// Amount returns the amount in the smallest currency unit
func (p *Price) Amount() int {
	return p.amount
}

// Currency returns the currency
func (p *Price) Currency() Currency {
	return p.currency
}

// DisplayAmount returns the amount in the standard currency unit
// For example, if amount is stored in cents (USD), it returns dollars
func (p *Price) DisplayAmount() float64 {
	switch p.currency {
	case USD, EUR:
		return float64(p.amount) / 100 // cents to dollars/euros
	default:
		return float64(p.amount) // KRW has no subdivision
	}
}

// String returns the formatted price with currency
func (p *Price) String() string {
	switch p.currency {
	case KRW:
		return fmt.Sprintf("₩%d", p.amount)
	case USD:
		return fmt.Sprintf("$%.2f", p.DisplayAmount())
	case EUR:
		return fmt.Sprintf("€%.2f", p.DisplayAmount())
	default:
		return fmt.Sprintf("%d %s", p.amount, p.currency)
	}
}

// Add returns a new Price with the sum of two prices
func (p *Price) Add(other *Price) (*Price, error) {
	if p.currency != other.currency {
		return nil, fmt.Errorf("cannot add prices with different currencies: %s and %s",
			p.currency, other.currency)
	}

	return NewPrice(p.amount+other.amount, p.currency)
}

// Subtract returns a new Price with the difference of two prices
func (p *Price) Subtract(other *Price) (*Price, error) {
	if p.currency != other.currency {
		return nil, fmt.Errorf("cannot subtract prices with different currencies: %s and %s",
			p.currency, other.currency)
	}

	return NewPrice(p.amount-other.amount, p.currency)
}

// MultiplyByQuantity returns a new Price multiplied by a quantity
func (p *Price) MultiplyByQuantity(quantity int) (*Price, error) {
	if quantity < 0 {
		return nil, fmt.Errorf("quantity cannot be negative: %d", quantity)
	}

	return NewPrice(p.amount*quantity, p.currency)
}

// ApplyDiscount returns a new Price with the discount applied
func (p *Price) ApplyDiscount(discountPercentage float64) (*Price, error) {
	if discountPercentage < 0 || discountPercentage > 100 {
		return nil, fmt.Errorf("invalid discount percentage: %.2f", discountPercentage)
	}

	discountedAmount := int(math.Round(float64(p.amount) * (1 - discountPercentage/100)))
	return NewPrice(discountedAmount, p.currency)
}

// Equals checks if two prices are equal
func (p *Price) Equals(other *Price) bool {
	if other == nil {
		return false
	}
	return p.amount == other.amount && p.currency == other.currency
}

// IsZero checks if the price is zero
func (p *Price) IsZero() bool {
	return p.amount == 0
}

// IsGreaterThan checks if this price is greater than another price
func (p *Price) IsGreaterThan(other *Price) (bool, error) {
	if p.currency != other.currency {
		return false, fmt.Errorf("cannot compare prices with different currencies: %s and %s",
			p.currency, other.currency)
	}
	return p.amount > other.amount, nil
}

// IsLessThan checks if this price is less than another price
func (p *Price) IsLessThan(other *Price) (bool, error) {
	if p.currency != other.currency {
		return false, fmt.Errorf("cannot compare prices with different currencies: %s and %s",
			p.currency, other.currency)
	}
	return p.amount < other.amount, nil
}
