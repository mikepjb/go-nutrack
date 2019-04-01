package stats

import (
	"sort"

	"github.com/mikepjb/nutrition/src/nutrition"
)

type IngredientStat struct {
	Name     string
	FoodItem nutrition.FoodItem
	Amount   float32
	Price    float32
}

// return food items in the amount they were used for the order.
func FoodItemUse(orders []nutrition.Order) []IngredientStat {
	foodItems := map[nutrition.FoodItem]float32{}
	ingredients := []nutrition.Ingredient{}
	ingredientStats := []IngredientStat{}

	for _, o := range orders {
		ingredients = append(ingredients, o.Ingredients()...)
	}

	for _, i := range ingredients {
		foodItems[i.FoodItem] += i.Amount
	}

	for fi, a := range foodItems {
		in := nutrition.Ingredient{
			Name:     fi.Name,
			FoodItem: fi,
			Amount:   a,
		}
		ins := IngredientStat{
			Name:     fi.Name,
			FoodItem: fi,
			Amount:   a,
			Price:    in.Price(),
		}
		ingredientStats = append(ingredientStats, ins)
	}

	sort.Slice(ingredientStats, func(i, j int) bool {
		return ingredientStats[i].Price > ingredientStats[j].Price
	})

	return ingredientStats
}

func FoodItemsTotalValue(orders []nutrition.Order) float32 {
	usedItems := FoodItemUse(orders)
	var totalValue float32

	for _, i := range usedItems {
		totalValue += i.Price
	}

	return totalValue
}
