package stats

import (
	"fmt"
	"sort"

	"github.com/mikepjb/nutrack/src/nutrition"
)

type IngredientStat struct {
	Name     string
	FoodItem nutrition.FoodItem
	Amount   float32
	Price    float32
	Protein  float32
	Carbs    float32
	Fat      float32
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
			Protein:  fi.Protein,
			Carbs:    fi.Carbs,
			Fat:      fi.Fat,
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

func WeeklyNutrition(orders []nutrition.Order) nutrition.Nutrition {
	return orders[0].Nutrition()
}

func DailyNutrition(orders []nutrition.Order) nutrition.Nutrition {
	weeklyNutrition := WeeklyNutrition(orders)

	return nutrition.Nutrition{
		Energy:  weeklyNutrition.Energy / 7,
		Fat:     weeklyNutrition.Fat / 7,
		Sfat:    weeklyNutrition.Sfat / 7,
		Carbs:   weeklyNutrition.Carbs / 7,
		Sugars:  weeklyNutrition.Sugars / 7,
		Protein: weeklyNutrition.Protein / 7,
	}
}

func TargetDailyNutrition(weight float32) nutrition.Nutrition {
	var fat float32 = 80
	var carbs float32 = 140
	var protein float32 = weight * 1.6

	return nutrition.Nutrition{
		Energy:  (fat * 9) + (carbs * 4) + (protein * 4),
		Fat:     fat,
		Sfat:    20,
		Carbs:   carbs,
		Sugars:  20,
		Protein: protein,
	}
}

func MacroRatio(n nutrition.Nutrition) string {
	total := n.Protein + n.Carbs + n.Fat
	return fmt.Sprintf(
		"%.2f : %.2f : %.2f",
		n.Fat/total,
		n.Carbs/total,
		n.Protein/total,
	)
}
