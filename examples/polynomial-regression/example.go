package main

import (
	"fmt"
	"log"

	"github.com/antonsegerkvist/regression"
)

func main() {

	points := 10

	data := []regression.Point32{
		{X: 0, Y: -1},
		{X: 0, Y: 0},
		{X: 0, Y: 2},
		{X: 1, Y: 1},
		{X: 1, Y: 2},
		{X: 1, Y: 4},
		{X: 2, Y: 5},
		{X: 2, Y: 7},
		{X: 2, Y: 6},
		{X: 3, Y: 4},
		{X: 3, Y: -1},
		{X: 3, Y: 2},
	}

	regression32 := regression.NewPolynomialRegression32(10)
	err := regression32.Train(&data)
	if err != nil {
		log.Fatal(err)
	}

	x := []float32{}
	y := []float32{}
	for i := 0; i < points; i++ {
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
