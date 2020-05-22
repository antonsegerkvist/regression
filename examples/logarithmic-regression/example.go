package main

import (
	"fmt"

	"github.com/antonsegerkvist/regression"
)

func main() {

	points := 10

	data := []regression.Point32{
		{X: 1, Y: 0},
		{X: 1, Y: 1},
		{X: 1, Y: 2},
		{X: 2, Y: 1},
		{X: 2, Y: 2},
		{X: 2, Y: 0},
		{X: 3, Y: 2},
		{X: 3, Y: 3},
		{X: 3, Y: 4},
		{X: 4, Y: 3},
		{X: 4, Y: 4},
		{X: 4, Y: 5},
	}

	regression32 := regression.NewLogarithmicRegression32()
	regression32.Train(&data)

	x := []float32{}
	y := []float32{}
	for i := float32(0.5); i < float32(points); i += 1.0 {
		x = append(x, 4.0/float32(points)*float32(i))
		y = append(y, regression32.Predict(4.0/float32(points)*float32(i)))
	}

	for _, v := range x {
		fmt.Printf("%f ", v)
	}
	fmt.Println("")

	for _, v := range y {
		fmt.Printf("%f ", v)
	}
	fmt.Println("")

}
