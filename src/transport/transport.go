// takes go structs and embeds function results in them for JSON transport
// e.g calculates the price of a recipe and include it as a RecipeTransport
// type.
package transport

import (
	"github.com/mikepjb/nutrition/src/nutrition"
	"github.com/mikepjb/nutrition/src/stats"
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
	Orders      []nutrition.Order
	Recipes     []nutrition.Recipe
	Ingredients []nutrition.Ingredient
	FoodItems   []nutrition.FoodItem
	Stats       Stats
}
