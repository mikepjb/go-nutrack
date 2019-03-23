// ingests nutrient plans
package nutrition

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type Recipe struct {
	Name        string
	Ingredients []Ingredient
}

func (r Recipe) Price() float32 {
	var price float32
	for _, i := range r.Ingredients {
		price += i.FoodItem.Price * i.Ratio()
	}
	return price
}

// A FoodItem is something that can be bought at a store (currently this is
// Mike's local Tesco)
type FoodItem struct {
	Name    string
	Desc    string
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
	Name     string
	FoodItem FoodItem
	Amount   int
}

// the amount of ingredient relative to the FoodItem's original amount. For
// example oats FoodItem is 100g but the Ingredient amount may only be 70g.
func (i Ingredient) Ratio() float32 {
	return float32(i.Amount) / float32(i.FoodItem.Amount)
}

func (i Ingredient) Price() float32 {
	return i.FoodItem.Price * i.Ratio()
}

type Order struct {
	Name    string
	Recipes []Recipe
}

func (o Order) Price() float32 {
	var price float32
	for _, r := range o.Recipes {
		price += r.Price()
	}
	return price
}

func (o Order) Ingredients() []Ingredient {
	ingredients := []Ingredient{}
	for _, r := range o.Recipes {
		ingredients = append(ingredients, r.Ingredients...)
	}
	return ingredients
}

// Plan is a collection of information about food and drink that you are
// consuming. Orders are the recipes you plan to cook in a given time frame
// (default 2 weeks), Recipes are the combination of Ingredients into a
// consumable. Ingredients are the raw materials used.
type Plan struct {
	Orders      []Order
	Recipes     []Recipe
	Ingredients []Ingredient
	FoodItem    []FoodItem
}

type OrderJSON struct {
	Name    string   `json:"name"`
	Recipes []string `json:"recipes"`
}

type RecipeJSON struct {
	Name        string   `json:"name"`
	Ingredients []string `json:"ingredients"`
}

type IngredientJSON struct {
	Name     string `json:"name"`
	FoodItem string `json:"foodItem"`
	Amount   int    `json:"amount"`
}

// There is a seperate internal JSON specific type used when ingesting a
// nutrient plan as JSON. This is because the items included in the plan are
// referenced to each other by name to keep things DRY.
type planJSON struct {
	Orders      []OrderJSON      `json:"orders"`
	Recipes     []RecipeJSON     `json:"recipes"`
	Ingredients []IngredientJSON `json:"ingredients"`
	FoodItems   []FoodItem       `json:"foodItem"`
}

func FoodItemByName(fis []FoodItem, name string) (FoodItem, error) {
	for _, e := range fis {
		if e.Name == name {
			return e, nil
		}
	}
	return FoodItem{}, errors.New("food item " + name + " not found")
}

func IngredientByName(ingredients []Ingredient, name string) (Ingredient, error) {
	for _, e := range ingredients {
		if e.Name == name {
			return e, nil
		}
	}
	return Ingredient{}, errors.New("ingredient " + name + " not found")
}

func RecipeByName(recipes []Recipe, name string) (Recipe, error) {
	for _, r := range recipes {
		if r.Name == name {
			return r, nil
		}
	}
	return Recipe{}, errors.New("recipe " + name + " not found")
}

// Reading a nutrient plan must be done in order so that food-items come first,
// ingredients come second, recipes third and finally orders because they are
// embedded in one another.
func ReadFile(path string) ([]Order, []Recipe, []Ingredient, []FoodItem) {
	dfile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("could not read file from path: %v\n", err)
	}

	var planJSON planJSON
	err = json.Unmarshal(dfile, &planJSON)
	if err != nil {
		fmt.Printf("could not read file from path: %v\n", err)
	}

	foodItems := planJSON.FoodItems

	var ingredients []Ingredient

	for _, e := range planJSON.Ingredients {
		foodItem, err := FoodItemByName(foodItems, e.Name)

		if err != nil {
			fmt.Println(err)
		}

		ingredients = append(ingredients, Ingredient{
			Name:     e.Name,
			FoodItem: foodItem,
			Amount:   e.Amount,
		})
	}

	var recipes []Recipe

	for _, e := range planJSON.Recipes {
		if err != nil {
			fmt.Println(err)
		}

		var recipeIngredients []Ingredient

		for _, i := range e.Ingredients {
			ingredient, err := IngredientByName(ingredients, i)

			if err != nil {
				fmt.Println(err)
			}

			recipeIngredients = append(recipeIngredients, ingredient)
		}

		recipes = append(recipes, Recipe{
			Name:        e.Name,
			Ingredients: recipeIngredients,
		})
	}

	var orders []Order

	for _, o := range planJSON.Orders {
		var orderRecipes []Recipe

		for _, r := range o.Recipes {
			recipe, err := RecipeByName(recipes, r)

			if err != nil {
				fmt.Println(err)
			}

			orderRecipes = append(orderRecipes, recipe)
		}

		orders = append(orders, Order{
			Name:    o.Name,
			Recipes: orderRecipes,
		})
	}

	return orders, recipes, ingredients, foodItems
}
