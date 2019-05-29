package main

import (
	"fmt"
	"github.com/mikepjb/nutrack/src/pretty"
	"github.com/mikepjb/nutrack/src/ref"
	"github.com/mikepjb/nutrack/src/transport"
	"github.com/mikepjb/nutrack/src/web"
	"log"
	"os"
	"strings"
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
