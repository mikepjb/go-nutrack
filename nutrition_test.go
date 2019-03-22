package main

import (
	"reflect"
	"testing"
)

var oatsItem FoodItem
var blueberriesItem FoodItem

func init() {
	oatsItem = FoodItem{"oats", 1000, 2, 0, 0, 0, 0, 0, 0}
	blueberriesItem = FoodItem{"blueberries", 200, 2, 0, 0, 0, 0, 0, 0}
}

func TestCostOfPorridge(t *testing.T) {
	oats := Ingredient{oatsItem, 70}
	blueberries := Ingredient{blueberriesItem, 20}

	porridge := Recipe{[]Ingredient{oats}}

	if porridge.Price() != 28.571428 {
		t.Errorf("wrong price for porridge: %v\n", porridge.Price())
	}

	porridgeWithBlueberries := Recipe{[]Ingredient{oats, blueberries}}

	if porridgeWithBlueberries.Price() != 48.571426 {
		t.Errorf(
			"wrong price for porridge with blueberries: %v\n",
			porridgeWithBlueberries.Price())
	}
}

func TestCostOfSteak(t *testing.T) {
	steakItem := FoodItem{"Steak", 300, 5, 0, 0, 0, 0, 0, 0}
	steak := Ingredient{steakItem, 300}

	if steak.Price() != 5 {
		t.Errorf("wrong price for steak: %v\n", steak.Price())
	}
}

// given a list of meals:
//   - return a list of ingredients required
//   - total cost
//   - total nutrition (energy kcal, protein, carbs, fat)
func TestWeeklyOrder(t *testing.T) {
	oats := Ingredient{oatsItem, 70}
	blueberries := Ingredient{blueberriesItem, 60}

	porridgeWithBlueberries := Recipe{[]Ingredient{oats, blueberries}}

	order := Order{[]Recipe{porridgeWithBlueberries, porridgeWithBlueberries}}

	if order.Price() != 70.47619 {
		t.Errorf("wrong total for order: %v\n", order.Price())
	}

	expectedIngredients := []Ingredient{oats, blueberries, oats, blueberries}

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
