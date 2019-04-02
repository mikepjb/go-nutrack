// a web interface for nutrack
package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mikepjb/nutrition/src/ref"
	"github.com/mikepjb/nutrition/src/stats"
	"github.com/mikepjb/nutrition/src/transport"
)

func routes() *http.ServeMux {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("site"))
	mux.Handle("/", fs)

	mux.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
		// if r.Method == "POST" {
		jsonPath := "src/ref/test-plan.json"
		var weight float32 = 88
		orders, recipes, ingredients, foodItems := ref.ReadFile(jsonPath)

		transportRecipes := []transport.Recipe{}

		for _, r := range recipes {
			transportRecipes = append(transportRecipes, transport.Recipe{
				r,
				r.Nutrition(),
				r.Price(),
				r.Nutrition().RatioString(),
			})
		}

		transportOrders := []transport.Order{}

		for _, o := range orders {
			recipes := []transport.Recipe{}
			for _, r := range o.Recipes {
				recipes = append(recipes, transport.Recipe{
					r,
					r.Nutrition(),
					r.Price(),
					r.Nutrition().RatioString(),
				})
			}
			transportOrders = append(transportOrders, transport.Order{
				o.Name,
				recipes,
			})
		}

		// for now do not handle input and return test json result.
		// fmt.Fprintln(w, "Thanks!")
		transport := transport.Transport{
			Orders:      transportOrders,
			Recipes:     transportRecipes,
			Ingredients: ingredients,
			FoodItems:   foodItems,
			Stats: transport.Stats{
				FoodItemUse:             stats.FoodItemUse(orders),
				TotalPriceOfIngredients: stats.FoodItemsTotalValue(orders),
				WeeklyNutrition:         stats.WeeklyNutrition(orders),
				DailyNutrition:          stats.DailyNutrition(orders),
				TargetDailyNutrition:    stats.TargetDailyNutrition(weight),
				MacroRatio:              stats.MacroRatio(stats.WeeklyNutrition(orders)),
			},
		}

		json.NewEncoder(w).Encode(transport)
		// }
	})

	return mux
}

func Serve() {
	port := "8080"
	mux := routes()
	fmt.Printf("Starting serving on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
