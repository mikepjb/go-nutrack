package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mikepjb/nutrition/src/nutrition"
)

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("no args supplied")
		os.Exit(0)
	}

	jsonPath := strings.Join(os.Args[1:], "")
	orders, recipes, ingredients, foodItems := nutrition.ReadFile(jsonPath)

	fmt.Printf("Orders: %v\n", orders)
	fmt.Printf("Recipes: %v\n", recipes)
	fmt.Printf("Ingredients: %v\n", ingredients)
	fmt.Printf("FoodItems: %v\n", foodItems)

	var totalRecipePrices float32

	for _, r := range recipes {
		totalRecipePrices += r.Price()
	}

	fmt.Printf("Total Prices for Recipes: %v\n", totalRecipePrices)
}
