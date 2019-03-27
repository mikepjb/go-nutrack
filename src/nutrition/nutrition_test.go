package nutrition

import (
	"reflect"
	"testing"
)

var oatsItem FoodItem
var blueberriesItem FoodItem

func init() {
	oatsItem = FoodItem{"oats", "some oats", 1000, 2, Nutrition{0, 0, 0, 0, 0, 0}}
	blueberriesItem = FoodItem{"blueberries", "some blueberries", 200, 2, Nutrition{0, 0, 0, 0, 0, 0}}
}

func TestCostOfPorridge(t *testing.T) {
	oats := Ingredient{"oats", oatsItem, 70}
	blueberries := Ingredient{"blueberries", blueberriesItem, 20}

	porridge := Recipe{"porridge", []Ingredient{oats}}

	if porridge.Price() != 0.14 {
		t.Errorf("wrong price for porridge: %v\n", porridge.Price())
	}

	porridgeWithBlueberries := Recipe{
		"porridge with blueberries",
		[]Ingredient{oats, blueberries},
	}

	if porridgeWithBlueberries.Price() != 0.34 {
		t.Errorf(
			"wrong price for porridge with blueberries: %v\n",
			porridgeWithBlueberries.Price())
	}
}

func TestCostOfSteak(t *testing.T) {
	steakItem := FoodItem{"Steak", "delicious", 300, 5, Nutrition{0, 0, 0, 0, 0, 0}}
	steak := Ingredient{"Steak", steakItem, 300}

	if steak.Price() != 5 {
		t.Errorf("wrong price for steak: %v\n", steak.Price())
	}
}

// given a list of meals:
//   - return a list of ingredients required
//   - total cost
//   - total nutrition (energy kcal, protein, carbs, fat)
func TestWeeklyOrder(t *testing.T) {
	oats := Ingredient{"oats", oatsItem, 70}
	blueberries := Ingredient{"blueberries", blueberriesItem, 60}

	porridgeWithBlueberries := Recipe{"porridge with blueberries", []Ingredient{oats, blueberries}}

	order := Order{"morning", []Recipe{porridgeWithBlueberries, porridgeWithBlueberries}}

	if order.Price() != 1.48 {
		t.Errorf("wrong total for order: %v\n", order.Price())
	}

	expectedIngredients := []Ingredient{oats, blueberries, oats, blueberries}

	if !reflect.DeepEqual(order.Ingredients(), expectedIngredients) {
		t.Errorf(
			"ingredients list does not match expectedIngredients: %v\n",
			order.Ingredients())
	}
}
