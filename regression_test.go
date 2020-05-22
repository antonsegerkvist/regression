package regression_test

import (
	"math"
	"testing"

	"github.com/antonsegerkvist/regression"
)

const PRECISSION_EPS = 0.0001

func TestLinearRegression32(t *testing.T) {

	tests := [][]regression.Point32{
		[]regression.Point32{
			{X: 0, Y: 0},
			{X: 1, Y: 1},
		},
		[]regression.Point32{
			{X: 1, Y: 2},
			{X: 2, Y: 3},
		},
		[]regression.Point32{
			{X: 2, Y: 4},
			{X: 3, Y: 5},
		},
		[]regression.Point32{
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

	linearRegression := regression.NewLinearRegression32()

	for i, test := range tests {
		linearRegression.Train(&test)
		prediction := linearRegression.Predict(0)

		if math.Abs(float64(prediction-predictions[i])) >= PRECISSION_EPS {
			t.Errorf("Got prediction %f; expected prediction %f", prediction, predictions[i])
		}
	}

}

func TestLinearRegression64(t *testing.T) {

	tests := [][]regression.Point64{
		[]regression.Point64{
			{X: 0, Y: 0},
			{X: 1, Y: 1},
		},
		[]regression.Point64{
			{X: 1, Y: 2},
			{X: 2, Y: 3},
		},
		[]regression.Point64{
			{X: 2, Y: 4},
			{X: 3, Y: 5},
		},
		[]regression.Point64{
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

	linearRegression := regression.NewLinearRegression64()

	for i, test := range tests {
		linearRegression.Train(&test)
		prediction := linearRegression.Predict(0)

		if math.Abs(prediction-predictions[i]) >= PRECISSION_EPS {
			t.Errorf("Got prediction %f; expected prediction %f", prediction, predictions[i])
		}
	}

}

func TestExponentialRegression32(t *testing.T) {

	tests := [][]regression.Point32{
		[]regression.Point32{
			{X: 1, Y: 1},
			{X: 2, Y: 2},
		},
		[]regression.Point32{
			{X: 1, Y: 3},
			{X: 3, Y: 1},
			{X: 4, Y: 2},
		},
	}

	predictions := []float32{
		0.5,
		3.050905515,
	}

	exponentialRegression := regression.NewExponentialRegression32()

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

func TestExponentialRegression64(t *testing.T) {

	tests := [][]regression.Point64{
		[]regression.Point64{
			{X: 1, Y: 1},
			{X: 2, Y: 2},
		},
		[]regression.Point64{
			{X: 1, Y: 3},
			{X: 3, Y: 1},
			{X: 4, Y: 2},
		},
	}

	predictions := []float64{
		0.5,
		3.050905515,
	}

	exponentialRegression := regression.NewExponentialRegression64()

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

func TestPolynomialRegression32(t *testing.T) {

	testOrders := []int{
		2,
	}

	tests := [][]regression.Point32{
		[]regression.Point32{
			{X: 1, Y: 1},
			{X: 2, Y: 2},
			{X: 3, Y: 1},
		},
	}

	predictions := []float32{
		-2,
	}

	for i, test := range tests {
		polynomialRegression := regression.NewPolynomialRegression32(testOrders[i])

		err := polynomialRegression.Train(&test)
		if err != nil {
			t.Errorf("Got prediction error; expected prediction %f", predictions[i])
		}
		prediction := polynomialRegression.Predict(0)

		if math.Abs(float64(prediction-predictions[i])) >= PRECISSION_EPS {
			t.Errorf("Got prediction %f; expected prediction %f", prediction, predictions[i])
		}
	}

}

func TestPolynomialRegression64(t *testing.T) {

	testOrders := []int{
		2,
	}

	tests := [][]regression.Point64{
		[]regression.Point64{
			{X: 1, Y: 1},
			{X: 2, Y: 2},
			{X: 3, Y: 1},
		},
	}

	predictions := []float64{
		-2,
	}

	for i, test := range tests {
		polynomialRegression := regression.NewPolynomialRegression64(testOrders[i])

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
