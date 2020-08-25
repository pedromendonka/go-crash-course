package main

import "fmt"

// Maps are key value pairs like dictionaries or like JS objects
func main() {
	// declare nil map
	// var m = map[string]int
	// m := map[string]int
	// declare empty map
	emails := make(map[string]string)
	emails["Bob"] = "bob@gmail.com"
	emails["Rob"] = "rob@gmail.com"
	delete(emails, "Bob")
	fmt.Println(emails)
	// Declare map and values
	mobile := map[string]string{
		"Bob": "1234556789",
		"Rob": "2323423423",
	}
	fmt.Println(mobile)
}
