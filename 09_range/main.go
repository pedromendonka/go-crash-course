package main

import "fmt"

func main() {
	ids := []int{22, 43, 75, 93}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}
	// Loop through ids - NO INDEX
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// range with map
	mobile := map[string]string{
		"Bob": "1234556789",
		"Rob": "2323423423",
	}
	for k, v := range mobile {
		fmt.Printf("%s: %s\n", k, v)
	}
}
