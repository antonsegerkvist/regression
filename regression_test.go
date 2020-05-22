package regression_test

import (
	"math"
	"testing"

	"github.com/antonsegerkvist/regression"
)

const PRECISSION_EPS = 0.0001

func TestLinearRegression2D32(t *testing.T) {

	tests := [][]regression.Point2D32{
		[]regression.Point2D32{
			{X: 0, Y: 0},
			{X: 1, Y: 1},
		},
		[]regression.Point2D32{
			{X: 1, Y: 2},
			{X: 2, Y: 3},
		},
		[]regression.Point2D32{
			{X: 2, Y: 4},
			{X: 3, Y: 5},
		},
		[]regression.Point2D32{
			{X: 0, Y: 3},
			{X: 1, Y: 3},
			{X: 2, Y: 3},
			{X: 3, Y: 3},
		},
	}

	predictions := []float32{
		0,
		1,
		2,
		3,
	}

	linearRegression := regression.NewLinearRegression2D32()

	for i, test := range tests {
		linearRegression.Train(&test)
		prediction := linearRegression.Predict(0)

		if math.Abs(float64(prediction-predictions[i])) >= PRECISSION_EPS {
			t.Errorf("Got prediction %f; expected prediction %f", prediction, predictions[i])
		}
	}

}

func TestLinearRegression2D64(t *testing.T) {

	tests := [][]regression.Point2D64{
		[]regression.Point2D64{
			{X: 0, Y: 0},
			{X: 1, Y: 1},
		},
		[]regression.Point2D64{
			{X: 1, Y: 2},
			{X: 2, Y: 3},
		},
		[]regression.Point2D64{
			{X: 2, Y: 4},
			{X: 3, Y: 5},
		},
		[]regression.Point2D64{
			{X: 0, Y: 3},
			{X: 1, Y: 3},
			{X: 2, Y: 3},
			{X: 3, Y: 3},
		},
	}

	predictions := []float64{
		0,
		1,
		2,
		3,
	}

	linearRegression := regression.NewLinearRegression2D64()

	for i, test := range tests {
		linearRegression.Train(&test)
		prediction := linearRegression.Predict(0)

		if math.Abs(prediction-predictions[i]) >= PRECISSION_EPS {
			t.Errorf("Got prediction %f; expected prediction %f", prediction, predictions[i])
		}
	}

}

func TestExponentialRegression2D32(t *testing.T) {

	tests := [][]regression.Point2D32{
		[]regression.Point2D32{
			{X: 1, Y: 1},
			{X: 2, Y: 2},
		},
		[]regression.Point2D32{
			{X: 1, Y: 3},
			{X: 3, Y: 1},
			{X: 4, Y: 2},
		},
	}

	predictions := []float32{
		0.5,
		3.050905515,
	}

	exponentialRegression := regression.NewExponentialRegression2D32()

	for i, test := range tests {
		err := exponentialRegression.Train(&test)
		if err != nil {
			t.Errorf("Got prediction error; expected prediction %f", predictions[i])
		}
		prediction := exponentialRegression.Predict(0)

		if math.Abs(float64(prediction-predictions[i])) >= PRECISSION_EPS {
			t.Errorf("Got prediction %f; expected prediction %f", prediction, predictions[i])
		}
	}

}

func TestExponentialRegression2D64(t *testing.T) {

	tests := [][]regression.Point2D64{
		[]regression.Point2D64{
			{X: 1, Y: 1},
			{X: 2, Y: 2},
		},
		[]regression.Point2D64{
			{X: 1, Y: 3},
			{X: 3, Y: 1},
			{X: 4, Y: 2},
		},
	}

	predictions := []float64{
		0.5,
		3.050905515,
	}

	exponentialRegression := regression.NewExponentialRegression2D64()

	for i, test := range tests {
		err := exponentialRegression.Train(&test)
		if err != nil {
			t.Errorf("Got prediction error; expected prediction %f", predictions[i])
		}
		prediction := exponentialRegression.Predict(0)

		if math.Abs(prediction-predictions[i]) >= PRECISSION_EPS {
			t.Errorf("Got prediction %f; expected prediction %f", prediction, predictions[i])
		}
	}

}

func TestPolynomialRegression2D64(t *testing.T) {

	testOrders := []int{
		2,
	}

	tests := [][]regression.Point2D64{
		[]regression.Point2D64{
			{X: 1, Y: 1},
			{X: 2, Y: 2},
			{X: 3, Y: 1},
		},
	}

	predictions := []float64{
		0,
	}

	for i, test := range tests {
		polynomialRegression := regression.NewPolynomialRegression2D64(testOrders[i])

		err := polynomialRegression.Train(&test)
		if err != nil {
			t.Errorf("Got prediction error; expected prediction %f", predictions[i])
		}
		prediction := polynomialRegression.Predict(0)

		if math.Abs(prediction-predictions[i]) >= PRECISSION_EPS {
			t.Errorf("Got prediction %f; expected prediction %f", prediction, predictions[i])
		}
	}

}
