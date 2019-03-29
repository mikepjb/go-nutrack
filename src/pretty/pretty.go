// A library for (you guessed it) pretty printing information
package pretty

import (
	"fmt"
	"os"
	"text/tabwriter"

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

func PrintOrdersNutrition(orders []nutrition.Order) {
	weeklyNutrition := orders[0].Nutrition()
	fmt.Printf("Total Weekly Nutrition: %+v\n", weeklyNutrition)

	dailyNutrition := nutrition.Nutrition{
		Energy:  weeklyNutrition.Energy / 7,
		Fat:     weeklyNutrition.Fat / 7,
		Sfat:    weeklyNutrition.Sfat / 7,
		Carbs:   weeklyNutrition.Carbs / 7,
		Sugars:  weeklyNutrition.Sugars / 7,
		Protein: weeklyNutrition.Protein / 7,
	}

	fmt.Printf("Average Daily Nutrition: %+v\n", dailyNutrition)

	var weight float32 = 88 // my weight in kg

	var fat float32 = 95
	var carbs float32 = 200
	var protein float32 = weight * 1.4

	targetNutrition := nutrition.Nutrition{
		Energy:  (fat * 9) + (carbs * 4) + (protein * 4),
		Fat:     fat,
		Sfat:    20,
		Carbs:   carbs,
		Sugars:  20,
		Protein: protein,
	}

	fmt.Printf("Target Daily Nutrition: %+v\n", targetNutrition)
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
	w := new(tabwriter.Writer).Init(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Printf("FoodItems:\n==================\n")
	for _, i := range items {
		fmt.Fprintf(w, "%v\t£%.2f\n", i.Desc, i.Price)
	}
	w.Flush()
	fmt.Printf("\n")
}

// this has application logic and should be moved to nutrition
func PrintFoodItemsUsed(orders []nutrition.Order) {
	w := new(tabwriter.Writer).Init(os.Stdout, 0, 0, 2, ' ', 0)
	foodItems := map[nutrition.FoodItem]float32{}
	var totalPrice float32

	ingredients := []nutrition.Ingredient{}

	for _, o := range orders {
		ingredients = append(ingredients, o.Ingredients()...)
	}

	for _, i := range ingredients {
		foodItems[i.FoodItem] += i.Amount
	}

	fmt.Println("List of FoodItems:\n==================")
	for fi, a := range foodItems {
		t := nutrition.Ingredient{
			Name:     fi.Name,
			FoodItem: fi,
			Amount:   a,
		}
		totalPrice += t.Price()
		fmt.Fprintf(w, "%v\t%vg\t£%.2f\n", t.Name, t.Amount, t.Price())
	}

	w.Flush()

	fmt.Printf("\nTotal Price for Ingredients used in Orders: £%.2f\n", totalPrice)
}
