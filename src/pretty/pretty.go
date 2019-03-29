// A library for (you guessed it) pretty printing information
package pretty

import (
	"fmt"

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
		fmt.Printf("%v: %+v, £%.2f\n", r.Name, r.Nutrition(), r.Price())
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
		fmt.Printf("%v: £%.2f\n", i.Desc, i.Price)
	}
	fmt.Printf("\n")
}

// this has application logic and should be moved to nutrition
func PrintFoodItemsUsed(orders []nutrition.Order) {
	foodItems := map[nutrition.FoodItem]float32{}
	var totalPrice float32

	ingredients := []nutrition.Ingredient{}

	for _, o := range orders {
		ingredients = append(ingredients, o.Ingredients()...)
	}

	for _, i := range ingredients {
		foodItems[i.FoodItem] += i.Amount
	}

	fmt.Println("List of FoodItems:")
	for fi, a := range foodItems {
		t := nutrition.Ingredient{
			Name:     fi.Name,
			FoodItem: fi,
			Amount:   a,
		}
		totalPrice += t.Price()
		fmt.Printf("%v, %vg: £%.2f\n", t.Name, t.Amount, t.Price())
	}

	fmt.Printf("Total Price for Ingredients used in Orders: £%.2f\n", totalPrice)
}
