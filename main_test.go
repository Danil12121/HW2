package main

import (
	"HomeWork2/calculator"
	"HomeWork2/shapes"
	"errors"
	"fmt"
	"math"
	"os"
	"reflect"
	"testing"
)

func Test_Shapes(t *testing.T) {
	test_cases := []struct {
		id    int
		input string
		want  float64
	}{
		{
			id:    1,
			input: "test.txt",
			want:  2.8*9.4 + 8*6 + math.Pi*5.7*5.7,
		},
		{
			id:    2,
			input: "test2.txt",
			want:  math.Pi*6*6 + 1*100,
		},
	}
	for _, test := range test_cases {
		t.Run(test.input, func(t *testing.T) {
			t.Cleanup(func() {
				t.Logf("comleted test number %d", test.id)
			})
			fileName := test.input

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

			get := calculator.TotalArea(li)

			if get != test.want {
				t.Errorf("Incorrect Answer: get %f, but want %f", get, test.want)
			}

			errc := f.Close()
			if errc != nil {
				fmt.Println(errc)
				return
			}
		})
	}
}
