'use strict'

var nutrientJSON

function tableData(text) {
  var td = document.createElement("td")
  var text = document.createTextNode(text)
  td.appendChild(text)
  return td
}

// no apostrophe :(
// also, seperate orders for each day is a better use of this model.
function todaysMenu(orders) {
  var dayOfWeek = new Date().getDay()
  var menu = orders[0].Recipes.slice(dayOfWeek*3, (dayOfWeek*3)+3)
  var menuSection = document.getElementById("menu-section")
  menu.forEach(function(r) {
    var rdiv = document.createElement("div")

    var span = document.createElement("span")
    var text = document.createTextNode(r.Name)
    span.appendChild(text)
    rdiv.appendChild(span)

    menuSection.appendChild(rdiv)
  })
}

function drawFoodItems(foodItemsUsed, total) {
  var tbody = document.getElementById("foodItems")

  foodItemsUsed.forEach(function(fi) {
    var tr = document.createElement("tr")

    tr.appendChild(tableData(fi.Name))
    tr.appendChild(tableData(fi.Protein))
    tr.appendChild(tableData(fi.Carbs))
    tr.appendChild(tableData(fi.Fat))
    tr.appendChild(tableData(fi.Amount + "g"))
    tr.appendChild(tableData("£" + fi.Price.toFixed(2)))

    tbody.appendChild(tr)
  })

  var tr = document.createElement("tr")
  tr.appendChild(tableData("Total"))
  tr.appendChild(tableData("--"))
  tr.appendChild(tableData("--"))
  tr.appendChild(tableData("--"))
  tr.appendChild(tableData("--"))
  tr.appendChild(tableData("£" + total.toFixed(2)))
  tbody.appendChild(tr)
}

function drawOrders(orders) {
  orders.forEach(function(o) {
    var e = document.createElement("h5")
    var text = document.createTextNode(o.Name)
    e.appendChild(text)
    var odiv = document.getElementById("orders")
    odiv.appendChild(e)

    o.Recipes.forEach(function(r) {
      var e = document.createElement("div")
      var text = document.createTextNode(r.Name)
      e.appendChild(text)
      odiv.appendChild(e)
    })
  })
}

function drawRecipes(recipes) {
  recipes.forEach(function(r) {
    var e = document.createElement("div")
    var text = document.createTextNode(r.Name)
    e.appendChild(text)
    var rdiv = document.getElementById("recipes")
    rdiv.appendChild(e)
  })
}

function update(id, value) {
  document.getElementById(id).innerHTML = value
}

function loadNutritionJSON() {
  var xhr = new XMLHttpRequest();
  xhr.open("GET", "calculate", true);
  xhr.onload = function() {
    console.log("loading nutrient json");
    nutrientJSON = JSON.parse(this.responseText);

    drawFoodItems(nutrientJSON.Stats.FoodItemUse,
      nutrientJSON.Stats.TotalPriceOfIngredients)
    drawOrders(nutrientJSON.Orders)
    drawRecipes(nutrientJSON.Recipes)
    update("spend-total",
      "£"+nutrientJSON.Stats.TotalPriceOfIngredients.toFixed(2))
    update("nutrition-total",
      nutrientJSON.Stats.DailyNutrition.Energy.toFixed(0)+" / "+nutrientJSON.Stats.TargetDailyNutrition.Energy.toFixed(0))
    todaysMenu(nutrientJSON.Orders)
  }
  xhr.send();
}

loadNutritionJSON()