package main

import "fmt"

type Recipe struct{}

func (Recipe) Price() float32 {
	return 0.7
}

func main() {
	fmt.Println("hello")
}
