package calculator

import (
	"HomeWork2/shapes"
)

func TotalArea(elems []shapes.Shape) float64 {
	res := 0.0
	for _, elem := range elems {
		sq := shapes.CallSquare(elem)
		res += sq
	}
	return res
}
