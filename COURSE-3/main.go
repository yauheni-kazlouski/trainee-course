package main

import (
	"flag"
	"fmt"
	"math"
)

// Shape interface is implemented by Rectangle and Circle and provides func Area() that calculates an area of given shape
type Shape interface{
	Area() float64
}

// Rectangle struct stands for rectangle shape and implements Shape interface
type Rectangle struct {
	width  	float64
	height 	float64
}

// Returns new *Rectangle of given height and width
func NewRectangle(w, h float64) *Rectangle {
	return &Rectangle{w, h}
}

// implementation of Area method of Shape interface for Rectangle
func (r *Rectangle) Area() float64 {
	return float64(r.height * r.width)
}

// Circle struct stands for circle shape and implements Shape interface
type Circle struct {
	radius	float64
}

// Returns new *Circle of given radius
func NewCircle(r float64) *Circle {
	return &Circle{r}
}

// implementation of Area method of Shape interface for Circle
func (c *Circle) Area() float64 {
	return math.Pi * float64(c.radius) * float64(c.radius)
}

func main() {
	var shape string
	var width, height, radius float64

	flag.StringVar(&shape, "shape", "", "accepts ['rectangle'|'circle'], not null")
	flag.Float64Var(&width, "width", 0, "")
	flag.Float64Var(&height, "height", 0, "")
	flag.Float64Var(&radius, "radius", 0, "")

	// Parsing all the arguements given via CLI
	flag.Parse()

	switch shape {
	case "rectangle":
		if width <= 0 {
			fmt.Println("Error: Flag -width used with -shape='rectangle' must take values > 0")
			return
		}

		if height <= 0 {
			fmt.Println("Error: Flag -height used with -shape='rectangle' must take values > 0")
			return
		}

		rec := NewRectangle(width, height)
		fmt.Println(rec.Area())

	case "circle":
		if radius <= 0 {
			fmt.Println("Error: Flag -radius used with -shape='circle' must take values > 0")
			return
		}

		crl := NewCircle(radius)
		fmt.Println(crl.Area())

	default:
		// Flag -shape is essential, if it's not given an error is displayed
		fmt.Println("Error: Flag -shape must take values [ 'rectangle' | 'circle' ]")
		return
	}
}