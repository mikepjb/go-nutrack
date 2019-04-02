'use strict'

var nutrientJSON

String.prototype.capitalize = function() {
  return this.replace(/\w\S*/g, function(txt) {
    return txt.charAt(0).toUpperCase() + txt.substr(1).toLowerCase();
  });
}

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

    var div = document.createElement("div")
    var text = document.createTextNode(r.Name.capitalize())
    div.appendChild(text)
    div.classList.add("name")
    rdiv.appendChild(div)

    var div = document.createElement("div")
    var text = document.createTextNode(r.RatioString)
    div.appendChild(text)
    div.classList.add("ratio")
    rdiv.appendChild(div)

    var div = document.createElement("div")

    var span = document.createElement("span")
    var text = document.createTextNode("protein")
    span.appendChild(text)
    div.appendChild(span)

    var span = document.createElement("span")
    var text = document.createTextNode("carbs")
    span.appendChild(text)
    div.appendChild(span)

    var span = document.createElement("span")
    var text = document.createTextNode("fat")
    span.appendChild(text)
    div.appendChild(span)

    div.classList.add("ratio-desc")
    rdiv.appendChild(div)

    var div = document.createElement("div")
    var text = document.createTextNode(r.Nutrition.Energy.toFixed(0)+" kcal")
    div.appendChild(text)
    div.classList.add("energy")
    rdiv.appendChild(div)

    menuSection.appendChild(rdiv)
  })
}

function drawFoodItems(foodItemsUsed, total) {
  var tbody = document.getElementById("foodItems")

  foodItemsUsed.forEach(function(fi) {
    var tr = document.createElement("tr")

    tr.appendChild(tableData(fi.Name.capitalize()))
    tr.appendChild(tableData(fi.Protein+"g"))
    tr.appendChild(tableData(fi.Carbs+"g"))
    tr.appendChild(tableData(fi.Fat+"g"))
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
    var text = document.createTextNode(o.Name.capitalize())
    e.appendChild(text)
    var odiv = document.getElementById("orders")
    odiv.appendChild(e)

    o.Recipes.forEach(function(r) {
      var e = document.createElement("div")
      var text = document.createTextNode(r.Name.capitalize())
      e.appendChild(text)
      odiv.appendChild(e)
    })
  })
}

function drawRecipes(recipes) {
  var tbody = document.getElementById("recipes")

  recipes.forEach(function(r) {
    var tr = document.createElement("tr")

    tr.appendChild(tableData(r.Name.capitalize()))
    tr.appendChild(tableData(r.Nutrition.Protein.toFixed(2)+"g"))
    tr.appendChild(tableData(r.Nutrition.Carbs.toFixed(2)+"g"))
    tr.appendChild(tableData(r.Nutrition.Fat.toFixed(2)+"g"))
    tr.appendChild(tableData("£" + r.Price.toFixed(2)))

    tbody.appendChild(tr)
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
