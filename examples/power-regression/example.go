package main

import (
	"fmt"

	"github.com/antonsegerkvist/regression"
)

func main() {

	points := 100

	data := []regression.Point32{
		{X: 83, Y: 183},
		{X: 71, Y: 168},
		{X: 64, Y: 171},
		{X: 69, Y: 178},
		{X: 69, Y: 176},
		{X: 64, Y: 172},
		{X: 68, Y: 165},
		{X: 59, Y: 158},
		{X: 81, Y: 183},
		{X: 91, Y: 182},
		{X: 57, Y: 163},
		{X: 65, Y: 175},
		{X: 58, Y: 164},
		{X: 62, Y: 175},
	}

	regression32 := regression.NewPowerRegression32()
	regression32.Train(&data)

	x := []float32{}
	y := []float32{}
	for i := float32(57); i < 91.0; i += ((91 - 57) / float32(points)) {
		x = append(x, float32(i))
		y = append(y, regression32.Predict(i))
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
