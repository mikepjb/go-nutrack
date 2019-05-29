// takes go structs and embeds function results in them for JSON transport
// e.g calculates the price of a recipe and include it as a transport.Recipe
// type.
package transport

import (
	"encoding/json"
	"io"

	"github.com/mikepjb/nutrack/src/nutrition"
	"github.com/mikepjb/nutrack/src/ref"
	"github.com/mikepjb/nutrack/src/stats"
)

type Stats struct {
	FoodItemUse             []stats.IngredientStat
	TotalPriceOfIngredients float32
	WeeklyNutrition         nutrition.Nutrition
	DailyNutrition          nutrition.Nutrition
	MacroRatio              string
	TargetDailyNutrition    nutrition.Nutrition
}

// this struct houses the other types to pass to the web view
// better name + abstraction?
type Transport struct {
	Orders      []Order
	Recipes     []Recipe
	Ingredients []nutrition.Ingredient
	FoodItems   []nutrition.FoodItem
	Stats       Stats
}

// TODO: I think there should be an interface for these..?
type Recipe struct {
	nutrition.Recipe
	Nutrition   nutrition.Nutrition
	Price       float32
	RatioString string
}

type Order struct {
	Name    string
	Recipes []Recipe
}

func WriteTestPlan(w io.Writer) {
	jsonPath := "src/ref/test-plan.json"
	var weight float32 = 88
	orders, recipes, ingredients, foodItems := ref.ReadFile(jsonPath)

	var transportRecipes []Recipe

	for _, r := range recipes {
		transportRecipes = append(transportRecipes, Recipe{
			r,
			r.Nutrition(),
			r.Price(),
			r.Nutrition().RatioString(),
		})
	}

	var transportOrders []Order

	for _, o := range orders {
		var recipes []Recipe
		for _, r := range o.Recipes {
			recipes = append(recipes, Recipe{
				r,
				r.Nutrition(),
				r.Price(),
				r.Nutrition().RatioString(),
			})
		}
		transportOrders = append(transportOrders, Order{
			o.Name,
			recipes,
		})
	}

	// for now do not handle input and return test json result.
	// fmt.Fprintln(w, "Thanks!")
	transport := Transport{
		Orders:      transportOrders,
		Recipes:     transportRecipes,
		Ingredients: ingredients,
		FoodItems:   foodItems,
		Stats: Stats{
			FoodItemUse:             stats.FoodItemUse(orders),
			TotalPriceOfIngredients: stats.FoodItemsTotalValue(orders),
			WeeklyNutrition:         stats.WeeklyNutrition(orders),
			DailyNutrition:          stats.DailyNutrition(orders),
			TargetDailyNutrition:    stats.TargetDailyNutrition(weight),
			MacroRatio:              stats.MacroRatio(stats.WeeklyNutrition(orders)),
		},
	}

	json.NewEncoder(w).Encode(transport)
}
