package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCalculatesTheCostOfPorridge(t *testing.T) {
	oats := Ingredient{"oats", 0, 0.7, 0, 0, 0, 0, 0, 0}
	blueberries := Ingredient{"blueberries", 0, 0.2, 0, 0, 0, 0, 0, 0}

	porridge := Recipe{[]Ingredient{oats}}

	if porridge.Price() != 0.7 {
		t.Errorf("wrong price for porridge: %v\n", porridge.Price())
	}

	porridgeWithBlueberries := Recipe{[]Ingredient{oats, blueberries}}

	if porridgeWithBlueberries.Price() != 0.9 {
		t.Errorf(
			"wrong price for porridge with blueberries: %v\n",
			porridgeWithBlueberries.Price())
	}
}

// given a list of meals:
//   - return a list of ingredients required
//   - total cost
//   - total nutrition (energy kcal, protein, carbs, fat)
func TestWeeklyOrder(t *testing.T) {
	oats := Ingredient{"oats", 70, 0.7, 0, 0, 0, 0, 0, 0}
	blueberries := Ingredient{"blueberries", 60, 0.2, 0, 0, 0, 0, 0, 0}

	porridgeWithBlueberries := Recipe{[]Ingredient{oats, blueberries}}

	order := Order{[]Recipe{porridgeWithBlueberries, porridgeWithBlueberries}}

	if order.Price() != 1.8 {
		t.Errorf("wrong total for order: %v\n", order.Price())
	}

	expectedIngredients := map[string]int{"blueberries": 120, "oats": 140}
	actualIngredients := order.Ingredients()

	fmt.Println(actualIngredients)
	fmt.Println(expectedIngredients)

	if !reflect.DeepEqual(order.Ingredients(), expectedIngredients) {
		t.Errorf(
			"ingredients list does not match expectedIngredients: %v\n",
			order.Ingredients())
	}
}

// nutrition should be able to read information about orders, recipes and
// ingredients from a JSON file.
func TestReadData(t *testing.T) {
	orders, recipes, ingredients := readData("nutrient-plan.json")

	if len(orders) != 1 {
		t.Errorf("wrong number of orders: %v\n", len(orders))
	}

	if len(recipes) != 1 {
		t.Errorf("wrong number of recipes: %v\n", len(recipes))
	}

	if len(ingredients) != 1 {
		t.Errorf("wrong number of ingredients: %v\n", len(ingredients))
	}
}
