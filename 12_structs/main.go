package main

import (
	"fmt"
	"strconv"
)

// Person Struct
type Person struct {
	firstName, lastName string
	age                 int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello I'm " + p.firstName + " " + p.lastName + " and I'm " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

func main() {
	fmt.Println("Structs")
	// Init a person
	pedro := Person{
		firstName: "Pedro",
		lastName:  "Mendonça",
		age:       45,
	}
	// alternative declaration
	joao := Person{"Joao", "Silva", 46}
	joao.hasBirthday()
	fmt.Println(pedro.age)
	fmt.Println(joao.firstName)
	fmt.Println(joao.greet())
}
