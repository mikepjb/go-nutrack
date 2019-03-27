// ref contains reference types that describe the structure of our input JSON
package ref

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/mikepjb/nutrition/src/nutrition"
)

type OrderReference struct {
	Name    string   `json:"name"`
	Recipes []string `json:"recipes"`
}

type RecipeReference struct {
	Name        string   `json:"name"`
	Ingredients []string `json:"ingredients"`
}

type IngredientReference struct {
	Name     string  `json:"name"`
	FoodItem string  `json:"foodItem"`
	Amount   float32 `json:"amount"`
}

// There is a seperate internal JSON specific type used when ingesting a
// nutrient plan as JSON. This is because the items included in the plan are
// referenced to each other by name to keep things DRY.
type planReference struct {
	Orders      []OrderReference      `json:"orders"`
	Recipes     []RecipeReference     `json:"recipes"`
	Ingredients []IngredientReference `json:"ingredients"`
	FoodItems   []nutrition.FoodItem  `json:"foodItem"`
}

func FoodItemByName(fis []nutrition.FoodItem, name string) (nutrition.FoodItem, error) {
	for _, e := range fis {
		if e.Name == name {
			return e, nil
		}
	}
	return nutrition.FoodItem{}, errors.New("food item " + name + " not found")
}

func IngredientByName(ingredients []nutrition.Ingredient, name string) (nutrition.Ingredient, error) {
	for _, e := range ingredients {
		if e.Name == name {
			return e, nil
		}
	}
	return nutrition.Ingredient{}, errors.New("ingredient " + name + " not found")
}

func RecipeByName(recipes []nutrition.Recipe, name string) (nutrition.Recipe, error) {
	for _, r := range recipes {
		if r.Name == name {
			return r, nil
		}
	}
	return nutrition.Recipe{}, errors.New("recipe " + name + " not found")
}

// Reading a nutrient plan must be done in order so that food-items come first,
// ingredients come second, recipes third and finally orders because they are
// embedded in one another.
func ReadFile(path string) ([]nutrition.Order, []nutrition.Recipe, []nutrition.Ingredient, []nutrition.FoodItem) {
	dfile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("could not read file from path: %v\n", err)
	}

	var planReference planReference
	err = json.Unmarshal(dfile, &planReference)
	if err != nil {
		fmt.Printf("could not read file from path: %v\n", err)
	}

	foodItems := planReference.FoodItems

	var ingredients []nutrition.Ingredient

	for _, e := range planReference.Ingredients {
		foodItem, err := FoodItemByName(foodItems, e.Name)

		if err != nil {
			fmt.Println(err)
		}

		ingredients = append(ingredients, nutrition.Ingredient{
			Name:     e.Name,
			FoodItem: foodItem,
			Amount:   e.Amount,
		})
	}

	var recipes []nutrition.Recipe

	for _, e := range planReference.Recipes {
		if err != nil {
			fmt.Println(err)
		}

		var recipeIngredients []nutrition.Ingredient

		for _, i := range e.Ingredients {
			ingredient, err := IngredientByName(ingredients, i)

			if err != nil {
				fmt.Println(err)
			}

			recipeIngredients = append(recipeIngredients, ingredient)
		}

		recipes = append(recipes, nutrition.Recipe{
			Name:        e.Name,
			Ingredients: recipeIngredients,
		})
	}

	var orders []nutrition.Order

	for _, o := range planReference.Orders {
		var orderRecipes []nutrition.Recipe

		for _, r := range o.Recipes {
			recipe, err := RecipeByName(recipes, r)

			if err != nil {
				fmt.Println(err)
			}

			orderRecipes = append(orderRecipes, recipe)
		}

		orders = append(orders, nutrition.Order{
			Name:    o.Name,
			Recipes: orderRecipes,
		})
	}

	return orders, recipes, ingredients, foodItems
}
