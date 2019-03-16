package main

import (
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

// bad name.
func readData(path string) ([]Order, []Recipe, []Ingredient) {
	orders := []Order{}
	recipe := []Ingredient{}
	ingredients := []Recipe{}

	dfile := ioutil.ReadFile(path)

}

func main() {
	jsonPath := strings.Join(os.Args[1:], "")

	orders := []Order{}
	recipe := []Ingredient{}
	ingredients := []Recipe{}

	readData(jsonPath, orders, recipe, ingredients)
}
