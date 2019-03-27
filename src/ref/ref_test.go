package ref

import "testing"

// nutrition should be able to read information about orders, recipes and
// ingredients from a JSON file.
func TestReadData(t *testing.T) {
	orders, recipes, ingredients, foodItems := ReadFile("test-plan.json")

	if len(orders) != 1 {
		t.Errorf("wrong number of orders: %v\n", len(orders))
	}

	if len(recipes) != 9 {
		t.Errorf("wrong number of recipes: %v\n", len(recipes))
	}

	if len(ingredients) != 6 {
		t.Errorf("wrong number of ingredients: %v\n", len(ingredients))
	}

	if len(foodItems) != 6 {
		t.Errorf("wrong number of food items: %v\n", len(foodItems))
	}
}
