package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mikepjb/nutrition/src/pretty"
	"github.com/mikepjb/nutrition/src/ref"
	"github.com/mikepjb/nutrition/src/transport"
	"github.com/mikepjb/nutrition/src/web"
)

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("no args supplied")
		os.Exit(0)
	}

	if os.Args[1] == "web" {
		web.Serve()
	} else if os.Args[1] == "generate" {
		fmt.Println("writing test-plan.json to site folder")
		f, err := os.Create("./site/test-plan.json")

		if err != nil {
			fmt.Errorf("problem creating test-plan.json: %v\n")
		}

		transport.WriteTestPlan(f)
		f.Close()
	} else {
		jsonPath := strings.Join(os.Args[1:], "")
		orders, recipes, ingredients, foodItems := ref.ReadFile(jsonPath)

		pretty.PrintOrders(orders)
		pretty.PrintOrdersNutrition(orders)
		pretty.PrintRecipes(recipes)
		pretty.PrintIngredients(ingredients)
		pretty.PrintFoodItems(foodItems)
		pretty.PrintFoodItemsUsed(orders)
	}
}
