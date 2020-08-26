package main

import (
	"fmt"
	"math"
)

// Shape Interface for shapes (methods for structures)
type Shape interface {
	area() float64
}

// Circle struct
type Circle struct {
	x, y, radius float64
}

// Rectangle structure
type Rectangle struct {
	w, h float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (r Rectangle) area() float64 {
	return r.w * r.h
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	fmt.Println("Interfaces")
	circle := Circle{0, 0, 5}
	var rectangle = Rectangle{5, 6}

	// Methods
	fmt.Println(circle.area())
	fmt.Println(rectangle.area())
	// With interfaces
	fmt.Println(getArea(circle))
	fmt.Println(getArea(rectangle))
}
