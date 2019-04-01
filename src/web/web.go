// a web interface for nutrack
package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mikepjb/nutrition/src/nutrition"
	"github.com/mikepjb/nutrition/src/ref"
)

func routes() *http.ServeMux {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("site"))
	mux.Handle("/", fs)

	mux.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
		// if r.Method == "POST" {
		jsonPath := "src/ref/test-plan.json"
		orders, recipes, ingredients, foodItems := ref.ReadFile(jsonPath)
		// for now do not handle input and return test json result.
		// fmt.Fprintln(w, "Thanks!")
		transport := nutrition.Transport{
			Orders:      orders,
			Recipes:     recipes,
			Ingredients: ingredients,
			FoodItems:   foodItems,
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
