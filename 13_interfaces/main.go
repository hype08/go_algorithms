package main

import (
	"fmt"
	"math"
)

// Shape interface.
type Shape interface {
	area() float64
}

// Circle struct.
type Circle struct {
	radius float64
}

// Rectangle struct.
type Rectangle struct {
	height, width float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle: %f\n", getArea(circle))
	fmt.Printf("Rectangle: %f\n", getArea(rectangle))
}
