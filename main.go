package main

import (
	"HomeWork2/calculator"
	"HomeWork2/shapes"
	"errors"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic(errors.New("Wrong name of file"))
	}
	fileName := os.Args[1]

	f, err := os.Open(fileName)
	if err != nil {
		panic(errors.New("Error with read file"))
	}

	var li []shapes.Shape
	for {
		var st string
		var fl1 float64
		var fl2 float64
		var n int
		n, err = fmt.Fscanln(f, &st, &fl1, &fl2)

		if n == 0 {
			break
		}
		if err != nil {
			panic(errors.New("error with parametrs (must be float)"))
		}

		if st == "Circle" {
			c1, er1 := shapes.CreateCircle(fl1)
			if er1 != nil {
				panic(er1)
			}
			li = append(li, c1)
		}
		if st == "Rectangle" {
			r1, er2 := shapes.CreateRectangle(fl1, fl2)
			if er2 != nil {
				panic(er2)
			}
			li = append(li, r1)
		}
	}

	fmt.Println(calculator.TotalArea(li))
}
