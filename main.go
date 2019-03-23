package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mikepjb/nutrition/src/nutrition"
)

func PrintOrders(orders []nutrition.Order) {
	var count int
	fmt.Printf("Orders:\n==================\n")
	for _, o := range orders {
		fmt.Printf("%v:\n", o.Name)
		for _, r := range o.Recipes {
			fmt.Printf("  %v\n", r.Name)
			count++
			if count == 3 {
				fmt.Printf("\n")
				count = 0
			}
		}
	}
}

func PrintRecipes(recipes []nutrition.Recipe) {
	fmt.Printf("Recipes:\n==================\n")
	for _, r := range recipes {
		fmt.Printf("%v: %v\n", r.Name, r.Price())
	}
	fmt.Printf("\n")
}

func PrintIngredients(ingredients []nutrition.Ingredient) {
	fmt.Printf("Ingredients:\n==================\n")
	for _, i := range ingredients {
		fmt.Printf("%v (%vg): %v\n", i.Name, i.Amount, i.Price())
	}
	fmt.Printf("\n")
}

func PrintFoodItems(items []nutrition.FoodItem) {
	fmt.Printf("FoodItems:\n==================\n")
	for _, i := range items {
		fmt.Printf("%v: %v\n", i.Desc, i.Price)
	}
	fmt.Printf("\n")
}

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("no args supplied")
		os.Exit(0)
	}

	jsonPath := strings.Join(os.Args[1:], "")
	orders, recipes, ingredients, foodItems := nutrition.ReadFile(jsonPath)

	PrintOrders(orders)
	PrintRecipes(recipes)
	PrintIngredients(ingredients)
	PrintFoodItems(foodItems)

	var totalRecipePrices float32

	for _, o := range orders {
		for _, r := range o.Recipes {
			totalRecipePrices += r.Price()
		}
	}

	fmt.Printf("Total Prices for Orders: %v\n", totalRecipePrices)
}
