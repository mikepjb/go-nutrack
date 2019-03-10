package main

import (
	"testing"
)

func TestCalculatesTheCostOfPorridge(t *testing.T) {
	oats := Ingredient{"oats", 0.7}
	blueberries := Ingredient{"blueberries", 0.2}

	porridge := Recipe{[]Ingredient{oats}}

	if porridge.Price() != 0.7 {
		t.Errorf("wrong price for porridge: %v\n", porridge.Price())
	}

	porridgeWithBlueberries := Recipe{[]Ingredient{oats, blueberries}}

	if porridgeWithBlueberries.Price() != 0.9 {
		t.Errorf(
			"wrong price for porridge with blueberries: %v\n",
			porridgeWithBlueberries.Price())
	}
}
