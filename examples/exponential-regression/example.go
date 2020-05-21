package main

import (
	"fmt"
	"log"

	"github.com/antonsegerkvist/regression"
)

func main() {

	points := 10

	data := []regression.Point2D32{
		{X: 4, Y: 20},
		{X: 4, Y: 30},
		{X: 4, Y: 24},
		{X: 1, Y: 1},
		{X: 1, Y: 2},
		{X: 1, Y: 4},
		{X: 2, Y: 5},
		{X: 2, Y: 7},
		{X: 2, Y: 6},
		{X: 3, Y: 7},
		{X: 3, Y: 10},
		{X: 3, Y: 12},
	}

	regression32 := regression.NewExponentialRegression2D32()
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
