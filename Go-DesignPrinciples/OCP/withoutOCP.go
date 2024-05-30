package OCP

import (
	"fmt"
)

type Food2 struct {
	Name  string
	Price int
}

func PrintPrice(food Food2, priceType string) {
	var finalPrice int
	switch priceType {
	case "Standard":
		finalPrice = food.Price
	case "Discount":
		finalPrice = food.Price - 10
	case "ExclusiveDiscount":
		finalPrice = food.Price - 20
	default:
		finalPrice = food.Price
	}

	fmt.Printf("%s %s Price: %d\n", food.Name, priceType, finalPrice)
}

// func main() {
// 	foods := []Food{
// 		{Name: "Burger", Price: 100},
// 		{Name: "Maggie", Price: 50},
// 	}

// 	for _, food := range foods {
// 		PrintPrice(food, "Standard")
// 		PrintPrice(food, "Discount")
// 		PrintPrice(food, "ExclusiveDiscount")
// 	}
// }
