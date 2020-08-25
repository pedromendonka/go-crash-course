package main

import "fmt"

func main() {
	x := 10
	y := 10

	// if else
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else if y < x {
		fmt.Printf("%d is less than %d\n", y, x)
	} else {
		fmt.Printf("%d is equal to %d\n", y, x)
	}

	// switch case
	color := "green"
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is not red or blue")
	}
}
