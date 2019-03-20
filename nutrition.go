package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Recipe struct {
	ingredients []Ingredient
}

func (r Recipe) Price() float32 {
	var price float32
	for _, i := range r.ingredients {
		price += i.Price()
	}
	return price
}

type Ingredient struct {
	name   string
	amount int     // amount in grams
	price  float32 // price per 100g
}

func (i Ingredient) Price() float32 {
	return i.price
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

func (o Order) Ingredients() map[string]int {
	ingredients := map[string]int{}
	for _, r := range o.recipes {
		for _, i := range r.ingredients {
			ingredients[i.name] += i.amount
		}
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

// bad name.
func readData(path string) ([]Order, []Recipe, []Ingredient) {
	// orders := []Order{}
	// recipes := []Recipe{}
	// ingredients := []Ingredient{}

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
