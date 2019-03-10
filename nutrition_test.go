package main

import (
	"fmt"
	"testing"
)

func TestCalculatesTheCostOfPorridge(t *testing.T) {
	// ingredients := []Ingredients{}}
	porridge := Recipe{}

	if porridge.Price() != 0.7 {
		fmt.Errorf("wrong price for porridge: %v\n", porridge.Price())
	}
}
