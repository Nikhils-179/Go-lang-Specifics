package main

import (
	"fmt"
	//"strings"

	"github.com/NikhilS-179/design-patterns/SRP"
	"github.com/NikhilS-179/design-patterns/OCP"
)

func main() {
	j := SRP.Journal{}
	j.AddEntries("I am adding entry")
	j.AddEntries("I am adding entry2")
	//fmt.Println(j.RemoveEntries("I am adding entry"))
	fmt.Println(j)
	//fmt.Println(strings.Join(j.Entries, "\n"))
	p := SRP.Persistance{"\n"}
	p.SaveToFile(&j, "journal.txt")

	//
	// Create burger
	burger := &OCP.Burger{}
	burger.Order("Cheese Burger", 1)

	// Create maggie
	maggie := &OCP.Maggie{}
	maggie.Order("Masala Maggie", 2)

	// Create price calculators
	standardPriceCalculator := &OCP.StandardPriceCalculator{}
	discountPriceCalculator := &OCP.DiscountPriceCalculator{}

	// Set price calculators for burger and maggie
	burger.SetPriceCalculator(standardPriceCalculator)
	maggie.SetPriceCalculator(discountPriceCalculator)

	// Calculate prices
	burger.CalculatePrice("50", 2) // Assuming price per unit is "50" and quantity is 2
	maggie.CalculatePrice("20", 3) // Assuming price per unit is "20" and quantity is 3
}
