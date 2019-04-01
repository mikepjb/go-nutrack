// ingests nutrient plans
package nutrition

// this struct houses the other types to pass to the web view
// better name + abstraction?
type Transport struct {
	Orders      []Order
	Recipes     []Recipe
	Ingredients []Ingredient
	FoodItems   []FoodItem
}

type Order struct {
	Name    string
	Recipes []Recipe
}

type Recipe struct {
	Name        string
	Ingredients []Ingredient
}

type Ingredient struct {
	Name     string
	FoodItem FoodItem
	Amount   float32
}

type Nutrition struct {
	Energy  float32 // kcal
	Fat     float32 // grams
	Sfat    float32 // saturated fat
	Carbs   float32 // total (incl. sugars) grams
	Sugars  float32 // grams
	Protein float32
}

type FoodItem struct { // items bought in at a store
	Name   string
	Desc   string
	Amount float32 // purchased amount in grams
	Price  float32 // retail price
	Nutrition
}

func (o Order) Nutrition() Nutrition {
	var totalNutrition Nutrition

	for _, r := range o.Recipes {
		rn := r.Nutrition()
		totalNutrition = Nutrition{
			Energy:  totalNutrition.Energy + rn.Energy,
			Fat:     totalNutrition.Fat + rn.Fat,
			Sfat:    totalNutrition.Sfat + rn.Sfat,
			Carbs:   totalNutrition.Carbs + rn.Carbs,
			Sugars:  totalNutrition.Sugars + rn.Sugars,
			Protein: totalNutrition.Protein + rn.Protein,
		}
	}

	return totalNutrition
}

func (r Recipe) Nutrition() Nutrition {
	var totalNutrition Nutrition

	for _, i := range r.Ingredients {
		totalNutrition = Nutrition{
			Energy:  totalNutrition.Energy + (i.FoodItem.Energy * i.NutritionRatio()),
			Fat:     totalNutrition.Fat + (i.FoodItem.Fat * i.NutritionRatio()),
			Sfat:    totalNutrition.Sfat + (i.FoodItem.Sfat * i.NutritionRatio()),
			Carbs:   totalNutrition.Carbs + (i.FoodItem.Carbs * i.NutritionRatio()),
			Sugars:  totalNutrition.Sugars + (i.FoodItem.Sugars * i.NutritionRatio()),
			Protein: totalNutrition.Protein + (i.FoodItem.Protein * i.NutritionRatio()),
		}
	}

	return totalNutrition
}

func (i Ingredient) Nutrition() Nutrition {
	return Nutrition{
		Energy:  i.FoodItem.Energy * i.NutritionRatio(),
		Fat:     i.FoodItem.Fat * i.NutritionRatio(),
		Sfat:    i.FoodItem.Sfat * i.NutritionRatio(),
		Carbs:   i.FoodItem.Carbs * i.NutritionRatio(),
		Sugars:  i.FoodItem.Sugars * i.NutritionRatio(),
		Protein: i.FoodItem.Protein * i.NutritionRatio(),
	}
}

func (r Recipe) Price() float32 {
	var price float32
	for _, i := range r.Ingredients {
		price += i.FoodItem.Price * i.WeightRatio()
	}
	return price
}

func (i Ingredient) Price() float32 {
	return i.FoodItem.Price * i.WeightRatio()
}

func (o Order) Price() float32 {
	var price float32
	for _, r := range o.Recipes {
		price += r.Price()
	}
	return price
}

// the amount of ingredient relative to the FoodItem's original amount. For
// example oats FoodItem is 100g but the Ingredient amount may only be 70g.
func (i Ingredient) WeightRatio() float32 {
	return i.Amount / i.FoodItem.Amount
}

func (i Ingredient) NutritionRatio() float32 {
	return i.Amount / 100
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
