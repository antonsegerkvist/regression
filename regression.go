package regression

import (
	"errors"
	"math"
)

/**********
 * ERRORS *
 **********/

//
// ErrUndeterminedSystem is returned when there are no solutions to
// the regression system.
//
var ErrUndeterminedSystem error = errors.New("The system is undetermined")

/************
 * ENTITIES *
 ************/

//
// Point2D32 contains coordinates in 2D euclidian space with 32 bit precision.
//
type Point2D32 struct {
	X float32
	Y float32
}

//
// Point2D64 contains coordinates in 2D euclidian space with 64 bit precision.
//
type Point2D64 struct {
	X float64
	Y float64
}

/**************
 * REGRESSION *
 **************/

/************************
 * LinearRegression2D32 *
 ************************/

//
// LinearRegression2D32 contains methods for performing linear regression on a
// two dimensional data set using 64 bit precission.
//
type LinearRegression2D32 struct {
	Order        int
	Coefficients []float32
}

//
// NewLinearRegression2D32 creates an instance LinearRegression2D64.
//
func NewLinearRegression2D32() *LinearRegression2D32 {
	return &LinearRegression2D32{
		Order:        2,
		Coefficients: []float32{0, 0},
	}
}

//
// Train calculates the coefficient vector from the provided data set.
//
func (lg *LinearRegression2D32) Train(points *[]Point2D32) error {

	if len(*points) < 2 {
		return ErrUndeterminedSystem
	}

	n := float32(len(*points))

	var ySum float32 = 0
	var xSum float32 = 0
	var x2Sum float32 = 0
	var xySum float32 = 0

	for _, v := range *points {
		ySum += v.Y
		xSum += v.X
		x2Sum += v.X * v.X
		xySum += v.X * v.Y
	}

	var denominator = n*x2Sum - xSum*xSum

	lg.Coefficients[0] = (ySum*x2Sum - xSum*xySum) / denominator
	lg.Coefficients[1] = (n*xySum - xSum*ySum) / denominator

	return nil

}

//
// Predict calculates the regression value at the given coordinate.
//
func (lg LinearRegression2D32) Predict(x float32) float32 {
	return lg.Coefficients[0] + lg.Coefficients[1]*x
}

/************************
 * LinearRegression2D64 *
 ************************/

//
// LinearRegression2D64 contains methods for performing linear regression on a
// two dimensional data set using 64 bit precission.
//
type LinearRegression2D64 struct {
	Order        int
	Coefficients []float64
}

//
// NewLinearRegression2D64 creates an instance LinearRegression2D64.
//
func NewLinearRegression2D64() *LinearRegression2D64 {
	return &LinearRegression2D64{
		Order:        2,
		Coefficients: []float64{0, 0},
	}
}

//
// Train calculates the coefficient vector from the provided data set.
//
func (lr *LinearRegression2D64) Train(points *[]Point2D64) error {

	if len(*points) < 2 {
		return ErrUndeterminedSystem
	}

	n := float64(len(*points))

	var ySum float64 = 0
	var xSum float64 = 0
	var x2Sum float64 = 0
	var xySum float64 = 0

	for _, v := range *points {
		ySum += v.Y
		xSum += v.X
		x2Sum += v.X * v.X
		xySum += v.X * v.Y
	}

	var denominator = n*x2Sum - xSum*xSum

	lr.Coefficients[0] = (ySum*x2Sum - xSum*xySum) / denominator
	lr.Coefficients[1] = (n*xySum - xSum*ySum) / denominator

	return nil

}

//
// Predict calculates the regression value at the given coordinate.
//
func (lr LinearRegression2D64) Predict(x float64) float64 {
	return lr.Coefficients[0] + lr.Coefficients[1]*x
}

/*****************************
 * ExponentialRegression2D32 *
 *****************************/

//
// ExponentialRegression2D32 contains methods for performing exponential regression on a
// two dimensional data set using 32 bit precission.
//
type ExponentialRegression2D32 struct {
	Coefficients []float32
}

//
// NewExponentialRegression2D32 creates an instance LinearRegression2D64.
//
func NewExponentialRegression2D32() *ExponentialRegression2D32 {
	return &ExponentialRegression2D32{
		Coefficients: []float32{0, 0},
	}
}

//
// Train calculates the coefficient vector from the provided data set.
//
func (er *ExponentialRegression2D32) Train(points *[]Point2D32) error {

	if len(*points) < 2 {
		return ErrUndeterminedSystem
	}

	n := float32(len(*points))

	var ySum float32 = 0
	var xSum float32 = 0
	var x2Sum float32 = 0
	var xySum float32 = 0

	for _, v := range *points {
		ySum += float32(math.Log(float64(v.Y)))
		xSum += v.X
		x2Sum += v.X * v.X
		xySum += v.X * float32(math.Log(float64(v.Y)))
	}

	var denominator = n*x2Sum - xSum*xSum

	er.Coefficients[0] = float32(math.Exp(float64((ySum*x2Sum - xSum*xySum) / denominator)))
	er.Coefficients[1] = (n*xySum - xSum*ySum) / denominator

	return nil

}

//
// Predict calculates the regression value at the given coordinate.
//
func (er ExponentialRegression2D32) Predict(x float32) float32 {
	return er.Coefficients[0] * float32(math.Exp(float64(er.Coefficients[1]*x)))
}

/*****************************
 * ExponentialRegression2D64 *
 *****************************/

//
// ExponentialRegression2D64 contains methods for performing exponential regression on a
// two dimensional data set using 64 bit precission.
//
type ExponentialRegression2D64 struct {
	Coefficients []float64
}

//
// NewExponentialRegression2D64 creates an instance LinearRegression2D64.
//
func NewExponentialRegression2D64() *ExponentialRegression2D64 {
	return &ExponentialRegression2D64{
		Coefficients: []float64{0, 0},
	}
}

//
// Train calculates the coefficient vector from the provided data set.
//
func (er *ExponentialRegression2D64) Train(points *[]Point2D64) error {

	if len(*points) < 2 {
		return ErrUndeterminedSystem
	}

	n := float64(len(*points))

	var ySum float64 = 0
	var xSum float64 = 0
	var x2Sum float64 = 0
	var xySum float64 = 0

	for _, v := range *points {
		ySum += math.Log(v.Y)
		xSum += v.X
		x2Sum += v.X * v.X
		xySum += v.X * math.Log(v.Y)
	}

	var denominator = n*x2Sum - xSum*xSum

	er.Coefficients[0] = math.Exp((ySum*x2Sum - xSum*xySum) / denominator)
	er.Coefficients[1] = (n*xySum - xSum*ySum) / denominator

	return nil

}

//
// Predict calculates the regression value at the given coordinate.
//
func (er ExponentialRegression2D64) Predict(x float64) float64 {
	return er.Coefficients[0] * math.Exp(er.Coefficients[1]*x)
}
