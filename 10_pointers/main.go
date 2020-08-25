package main

import "fmt"

func main() {
	a := 5
	b := &a
	fmt.Println(a, b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Println(a, *b)

	// change value with a pointer
	*b = 10
	fmt.Println(a)
}
