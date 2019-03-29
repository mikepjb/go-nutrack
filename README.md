# nutrition

An implementation of my food calculator in Golang.

It works out how much nutrition you get for a given meal and how much it costs.
It will also aggregate that across a week and tell you what ingredients you'll
need to buy.

### TODO

Include vitamins and minerals e.g chives have little macronutritional value but
a tablespoon provides vitamin K, calcium, magnesium etc..

### Domain Qualities After Building

This is a section to collect things I missed when modelling the domain at the
beginning.

- FoodItems expire
- FoodItems can be broken down into sub FoodItems (e.g a whole chicken can be
  made into 2 thighs, 2 breasts, 2 wings, stock OR cooked whole)
- Recipes can be divided into Portions - I am using Recipes as if they were
  single portions.. which works okay

### Notes on Nutrition

Weight loss - find your metabolic rate and set a 500 calorie deficit.

RDA for Protein based on:
  - Kent University study - 1.4 grams per kg of bodyweight for strength training.
  - Susan M Kleiner from Case Western Reserve University - 1.6-2.2 grams per kg.
  - Personal experience - 1.6 is pretty good for moderate exercise (3 sessions)
