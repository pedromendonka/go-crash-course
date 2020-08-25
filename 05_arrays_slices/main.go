package main

import "fmt"

func main() {
	// Arrays (fixed size)
	var fruits [2]string
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fmt.Println(fruits)
	otherFruits := [2]string{"Pineaple", "Grape"}
	fmt.Println(otherFruits)
	// Slices
	fruitSlice := []string{"Apple", "Cherry", "Banana"}
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
}
