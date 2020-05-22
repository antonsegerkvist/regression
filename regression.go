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
// NewLinearRegression2D32 creates an instance LinearRegression2D32.
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
// ExponentialRegression2D32 contains methods for performing exponential
// regression on a two dimensional data set using 32 bit precission.
//
type ExponentialRegression2D32 struct {
	Coefficients []float32
}

//
// NewExponentialRegression2D32 creates an instance ExponentialRegression2D32.
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
// NewExponentialRegression2D64 creates an instance ExponentialRegression2D64.
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

/****************************
 * PolynomialRegression2D32 *
 ****************************/

//
// PolynomialRegression2D32 contains methods for performing polynomial regression on a
// two dimensional data set using 32 bit precission.
//
type PolynomialRegression2D32 struct {
	Order        int
	Coefficients []float32
}

//
// NewPolynomialRegression2D32 creates an instance PolynomialRegression2D32.
//
func NewPolynomialRegression2D32(order int) *PolynomialRegression2D32 {

	coefficients := []float32{}
	for i := 0; i <= order; i++ {
		coefficients = append(coefficients, 0)
	}

	return &PolynomialRegression2D32{
		Order:        order,
		Coefficients: coefficients,
	}

}

//
// Train calculates the coefficient vector from the provided data set.
//
func (pr *PolynomialRegression2D32) Train(points *[]Point2D32) error {

	for i := range pr.Coefficients {
		pr.Coefficients[i] = 0
	}

	matrix := [][]float32{}
	for i := 0; i <= pr.Order; i++ {
		matrix = append(matrix, []float32{})
		for j := 0; j <= pr.Order; j++ {
			matrix[i] = append(matrix[i], 0)
		}
	}

	for i := 0; i <= pr.Order; i++ {
		for j := i; j <= pr.Order; j++ {
			var tmp float32 = 0

			for _, v := range *points {
				tmp += float32(math.Pow(float64(v.X), float64(i+j)))
			}

			matrix[i][j] = tmp
			matrix[j][i] = tmp
		}
	}

	vector := []float32{}
	for i := 0; i <= pr.Order; i++ {
		var tmp float32 = 0

		for _, v := range *points {
			tmp += v.Y * float32(math.Pow(float64(v.X), float64(i)))
		}

		vector = append(vector, tmp)
	}

	var h int = 0
	var k int = 0

	for h <= pr.Order && k <= pr.Order {
		maxRowIndex := 0
		var maxRowValue float32 = 0

		for i := h; i <= pr.Order; i++ {
			if maxRowValue < float32(math.Abs(float64(matrix[i][k]))) {
				maxRowIndex = i
				maxRowValue = float32(math.Abs(float64(matrix[i][k])))
			}
		}

		if maxRowValue == 0 {
			return ErrUndeterminedSystem
		}

		temp1 := vector[h]
		vector[h] = vector[maxRowIndex]
		vector[maxRowIndex] = temp1

		for i := k; i <= pr.Order; i++ {
			temp2 := matrix[h][i]
			matrix[h][i] = matrix[maxRowIndex][i]
			matrix[maxRowIndex][i] = temp2
		}

		for i := h + 1; i <= pr.Order; i++ {
			f := matrix[i][k] / matrix[h][k]
			matrix[i][k] = 0
			vector[i] = vector[i] - vector[h]*f
			for j := k + 1; j <= pr.Order; j++ {
				matrix[i][j] = matrix[i][j] - matrix[h][j]*f
			}
		}

		h++
		k++
	}

	for i := pr.Order; i >= 0; i-- {
		pr.Coefficients[i] = vector[i]
		for j := i + 1; j <= pr.Order; j++ {
			pr.Coefficients[i] -= matrix[i][j] * pr.Coefficients[j]
		}
		pr.Coefficients[i] /= matrix[i][i]
	}

	return nil

}

//
// Predict calculates the regression value at the given coordinate.
//
func (pr PolynomialRegression2D32) Predict(x float32) float32 {
	var ret float32 = 0
	for i, v := range pr.Coefficients {
		ret += float32(math.Pow(float64(x), float64(i))) * v
	}
	return ret
}

/****************************
 * PolynomialRegression2D64 *
 ****************************/

//
// PolynomialRegression2D64 contains methods for performing polynomial regression on a
// two dimensional data set using 64 bit precission.
//
type PolynomialRegression2D64 struct {
	Order        int
	Coefficients []float64
}

//
// NewPolynomialRegression2D64 creates an instance PolynomialRegression2D64.
//
func NewPolynomialRegression2D64(order int) *PolynomialRegression2D64 {

	coefficients := []float64{}
	for i := 0; i <= order; i++ {
		coefficients = append(coefficients, 0)
	}

	return &PolynomialRegression2D64{
		Order:        order,
		Coefficients: coefficients,
	}

}

//
// Train calculates the coefficient vector from the provided data set.
//
func (pr *PolynomialRegression2D64) Train(points *[]Point2D64) error {

	for i := range pr.Coefficients {
		pr.Coefficients[i] = 0
	}

	matrix := [][]float64{}
	for i := 0; i <= pr.Order; i++ {
		matrix = append(matrix, []float64{})
		for j := 0; j <= pr.Order; j++ {
			matrix[i] = append(matrix[i], 0)
		}
	}

	for i := 0; i <= pr.Order; i++ {
		for j := i; j <= pr.Order; j++ {
			var tmp float64 = 0

			for _, v := range *points {
				tmp += math.Pow(v.X, float64(i+j))
			}

			matrix[i][j] = tmp
			matrix[j][i] = tmp
		}
	}

	vector := []float64{}
	for i := 0; i <= pr.Order; i++ {
		var tmp float64 = 0

		for _, v := range *points {
			tmp += v.Y * math.Pow(v.X, float64(i))
		}

		vector = append(vector, tmp)
	}

	var h int = 0
	var k int = 0

	for h <= pr.Order && k <= pr.Order {
		maxRowIndex := 0
		var maxRowValue float64 = 0

		for i := h; i <= pr.Order; i++ {
			if maxRowValue < math.Abs(matrix[i][k]) {
				maxRowIndex = i
				maxRowValue = math.Abs(matrix[i][k])
			}
		}

		if maxRowValue == 0 {
			return ErrUndeterminedSystem
		}

		temp1 := vector[h]
		vector[h] = vector[maxRowIndex]
		vector[maxRowIndex] = temp1

		for i := k; i <= pr.Order; i++ {
			temp2 := matrix[h][i]
			matrix[h][i] = matrix[maxRowIndex][i]
			matrix[maxRowIndex][i] = temp2
		}

		for i := h + 1; i <= pr.Order; i++ {
			f := matrix[i][k] / matrix[h][k]
			matrix[i][k] = 0
			vector[i] = vector[i] - vector[h]*f
			for j := k + 1; j <= pr.Order; j++ {
				matrix[i][j] = matrix[i][j] - matrix[h][j]*f
			}
		}

		h++
		k++
	}

	for i := pr.Order; i >= 0; i-- {
		pr.Coefficients[i] = vector[i]
		for j := i + 1; j <= pr.Order; j++ {
			pr.Coefficients[i] -= matrix[i][j] * pr.Coefficients[j]
		}
		pr.Coefficients[i] /= matrix[i][i]
	}

	return nil

}

//
// Predict calculates the regression value at the given coordinate.
//
func (pr PolynomialRegression2D64) Predict(x float64) float64 {
	var ret float64 = 0
	for i, v := range pr.Coefficients {
		ret += math.Pow(x, float64(i)) * v
	}
	return ret
}
