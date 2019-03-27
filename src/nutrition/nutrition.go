// ingests nutrient plans
package nutrition

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

type FoodItem struct { // items bought in at a store
	Name    string
	Desc    string
	Amount  float32 // purchased amount in grams
	Price   float32 // retail price
	Energy  float32 // kcal
	Fat     float32 // grams
	Sfat    float32 // saturated fat
	Carbs   float32 // total (incl. sugars) grams
	Sugars  float32 // grams
	Protein float32
}

type Ingredient struct {
	Name     string
	FoodItem FoodItem
	Amount   float32
}

// the amount of ingredient relative to the FoodItem's original amount. For
// example oats FoodItem is 100g but the Ingredient amount may only be 70g.
func (i Ingredient) Ratio() float32 {
	return i.Amount / i.FoodItem.Amount
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
