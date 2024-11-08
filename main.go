package main

import (
	"HomeWork2/calculator"
	"HomeWork2/shapes"
	"errors"
	"fmt"
	"os"
	"reflect"
)

const (
	CIRCLE    = "Circle"
	RECTANGLE = "Rectangle"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(errors.New("Wrong name of file"))
		return
	}
	fileName := os.Args[1]

	f, err := os.Open(fileName)

	if err != nil {
		fmt.Println(errors.New("Error with read file"))
		return
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
			fmt.Println(errors.New("error with parametrs (must be float)"))
			return
		}

		if st == CIRCLE {
			var m float64
			if reflect.TypeOf(fl1) != reflect.TypeOf(m) {
				fmt.Println(errors.New("radius must be float64"))
				return
			}
			c1 := shapes.CreateCircle(fl1)
			li = append(li, c1)
		}
		if st == RECTANGLE {
			var m, n float64
			if reflect.TypeOf(fl1) != reflect.TypeOf(m) {
				fmt.Println(errors.New("width must be float64"))
				return
			}
			if reflect.TypeOf(fl2) != reflect.TypeOf(n) {
				fmt.Println(errors.New("length must be float64"))
				return
			}
			r1 := shapes.CreateRectangle(fl1, fl2)
			li = append(li, r1)
		}
	}

	fmt.Println(calculator.TotalArea(li))

	errc := f.Close()
	if errc != nil {
		fmt.Println(errc)
		return
	}
}
