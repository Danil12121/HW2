package shapes

import (
	"errors"
	"math"
	"reflect"
)

type Shape interface {
	Square() float64
}
type Circle struct {
	radius float64
}

type Rectangle struct {
	length, width float64
}

func CreateCircle(radius any) (Circle, error) {
	var m float64
	if reflect.TypeOf(radius) != reflect.TypeOf(m) {
		return Circle{}, errors.New("radius must be float64")
	}
	m = radius.(float64)
	return Circle{radius: m}, nil
}

func CreateRectangle(width, length any) (Rectangle, error) {
	var m, n float64
	if reflect.TypeOf(width) != reflect.TypeOf(m) {
		return Rectangle{}, errors.New("width must be float64")
	}
	if reflect.TypeOf(length) != reflect.TypeOf(n) {
		return Rectangle{}, errors.New("length must be float64")
	}
	m = width.(float64)
	n = length.(float64)
	return Rectangle{width: m, length: n}, nil
}

func (c Circle) Square() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) Square() float64 {
	return r.width * r.length
}

func CallSquare(s Shape) float64 {
	return s.Square()
}
