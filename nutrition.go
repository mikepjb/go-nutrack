package main

import "fmt"

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
	name  string
	price float32 // price per 100g
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

func main() {
	fmt.Println("hello")
}
