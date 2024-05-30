package OCP

import (
	"fmt"
	"strconv"
)

// Food interface defines methods for ordering food and setting price calculator.
type Food interface {
	Order(name string, foodType int)
	SetPriceCalculator(pc PriceCalculator)
	CalculatePrice(pricePerUnit string, quantity int)
}

// PriceCalculator interface defines method for calculating price.
type PriceCalculator interface {
	Price(foodType string, pricePerUnit string, quantity int)
}

// Burger struct representing a burger food item.
type Burger struct {
	Name            string
	Type            int
	priceCalculator PriceCalculator
}

// Order orders a burger.
func (b *Burger) Order(name string, foodType int) {
	b.Name = name
	b.Type = foodType
}

// SetPriceCalculator sets the price calculator for the burger.
func (b *Burger) SetPriceCalculator(pc PriceCalculator) {
	b.priceCalculator = pc
}

// CalculatePrice calculates the price of the burger.
func (b *Burger) CalculatePrice(pricePerUnit string, quantity int) {
	b.priceCalculator.Price(b.Name, pricePerUnit, quantity)
}

// Maggie struct representing a maggie food item.
type Maggie struct {
	Name            string
	Type            int
	priceCalculator PriceCalculator
}

// Order orders a maggie.
func (m *Maggie) Order(name string, foodType int) {
	m.Name = name
	m.Type = foodType
}

// SetPriceCalculator sets the price calculator for the maggie.
func (m *Maggie) SetPriceCalculator(pc PriceCalculator) {
	m.priceCalculator = pc
}

// CalculatePrice calculates the price of the maggie.
func (m *Maggie) CalculatePrice(pricePerUnit string, quantity int) {
	m.priceCalculator.Price(m.Name, pricePerUnit, quantity)
}

// StandardPriceCalculator struct for calculating standard price.
type StandardPriceCalculator struct{}

// Price calculates the standard price.
func (s *StandardPriceCalculator) Price(foodType string, pricePerUnit string, quantity int) {
	price, _ := strconv.Atoi(pricePerUnit)
	totalPrice := price * quantity * 100
	fmt.Printf("%s Standard Price: %d\n", foodType, totalPrice)
}

// DiscountPriceCalculator struct for calculating discounted price.
type DiscountPriceCalculator struct{}

// Price calculates the discounted price.
func (d *DiscountPriceCalculator) Price(foodType string, pricePerUnit string, quantity int) {
	price, _ := strconv.Atoi(pricePerUnit)
	totalPrice := price * quantity * 100
	discount := totalPrice * 10 / 100 // 10% discount
	finalPrice := totalPrice - discount
	fmt.Printf("%s Discounted Price: %d\n", foodType, finalPrice)
}

func main() {
	// Create burger
	burger := &Burger{}
	burger.Order("Cheese Burger", 1)

	// Create maggie
	maggie := &Maggie{}
	maggie.Order("Masala Maggie", 2)

	// Create price calculators
	standardPriceCalculator := &StandardPriceCalculator{}
	discountPriceCalculator := &DiscountPriceCalculator{}

	// Set price calculators for burger and maggie
	burger.SetPriceCalculator(standardPriceCalculator)
	maggie.SetPriceCalculator(discountPriceCalculator)

	// Calculate prices
	burger.CalculatePrice("50", 2) // Assuming price per unit is "50" and quantity is 2
	maggie.CalculatePrice("20", 3) // Assuming price per unit is "20" and quantity is 3
}
