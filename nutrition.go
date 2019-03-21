package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

type Recipe struct {
	ingredients []Ingredient
}

func (r Recipe) Price() float32 {
	var price float32
	for _, i := range r.ingredients {
		price += i.FoodItem.Price * i.Ratio()
	}
	// this casting is lunacy. surely there is a better way.
	return float32(math.Ceil(float64(price)))
}

// A FoodItem is something that can be bought at a store (currently this is
// Mike's local Tesco)
type FoodItem struct {
	Name    string
	Amount  int     // purchased amount in grams
	Price   float32 // retail price
	Energy  int     // kcal
	Fat     float32 // grams
	Sfat    float32 // saturated fat
	Carbs   float32 // total (incl. sugars) grams
	Sugars  float32 // grams
	Protein float32
}

type Ingredient struct {
	FoodItem FoodItem
	Amount   int
}

// the amount of ingredient relative to the FoodItem's original amount. For
// example oats FoodItem is 100g but the Ingredient amount may only be 70g.
func (i Ingredient) Ratio() float32 {
	return float32(i.FoodItem.Amount) / float32(i.Amount)
}

func (i Ingredient) Price() float32 {
	return i.FoodItem.Price * i.Ratio()
}

type Order struct {
	recipes []Recipe
}

func (o Order) Price() float32 {
	var price float32
	for _, r := range o.recipes {
		price += r.Price()
	}
	return price
}

func (o Order) Ingredients() []Ingredient {
	ingredients := []Ingredient{}
	for _, r := range o.recipes {
		ingredients = append(ingredients, r.ingredients...)
	}
	return ingredients
}

// NutrientPlan is a collection of information about food and drink that you are
// consuming. Orders are the recipes you plan to cook in a given time frame
// (default 2 weeks), Recipes are the combination of Ingredients into a
// consumable. Ingredients are the raw materials used.
type NutrientPlan struct {
	Orders      []Order      `json:"orders"`
	Recipes     []Recipe     `json:"recipes"`
	Ingredients []Ingredient `json:"ingredients"`
}

func readData(path string) ([]Order, []Recipe, []Ingredient) {
	dfile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("could not read file from path: %v\n", err)
	}

	var nutrientPlan NutrientPlan
	err = json.Unmarshal(dfile, &nutrientPlan)
	if err != nil {
		fmt.Printf("could not read file from path: %v\n", err)
	}

	// return orders, recipes, ingredients
	return nutrientPlan.Orders, nutrientPlan.Recipes, nutrientPlan.Ingredients

}

func main() {
	jsonPath := strings.Join(os.Args[1:], "")
	orders, recipes, ingredients := readData(jsonPath)

	fmt.Println(orders, recipes, ingredients)
}
