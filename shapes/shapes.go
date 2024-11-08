package shapes

import (
	"math"
)

type Shape interface {
	Area() float64
}
type Circle struct {
	radius float64
}

type Rectangle struct {
	length, width float64
}

func CreateCircle(radius float64) Circle {
	return Circle{radius: radius}
}

func CreateRectangle(width, length float64) Rectangle {
	return Rectangle{width: width, length: length}
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) Area() float64 {
	return r.width * r.length
}
