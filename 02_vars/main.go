package main

import "fmt"

func main() {
	var name string = "Pedro"
	// Infered
	var surname = "Mendonka"
	// shorthand inside function
	age := 45
	// multi
	eyes, hair := "brown", "black"

	fmt.Println(name, surname, age)
	fmt.Println("Eyes:", eyes, "and hair:", hair)
}
