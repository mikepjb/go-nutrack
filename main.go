package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mikepjb/nutrition/src/nutrition"
	"github.com/mikepjb/nutrition/src/pretty"
)

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("no args supplied")
		os.Exit(0)
	}

	jsonPath := strings.Join(os.Args[1:], "")
	orders, recipes, ingredients, foodItems := nutrition.ReadFile(jsonPath)

	pretty.PrintOrders(orders)
	pretty.PrintRecipes(recipes)
	pretty.PrintIngredients(ingredients)
	pretty.PrintFoodItems(foodItems)

	var totalRecipePrices float32

	for _, o := range orders {
		for _, r := range o.Recipes {
			totalRecipePrices += r.Price()
		}
	}

	fmt.Printf("Total Prices for Orders: %v\n", totalRecipePrices)
}
