
<div align="center">
<h1>regression</h1>
<br/>
<br/>
<p>
regression is a Go package containing a collection of linear least-squares fitting methods for simple data analysis.
</p>
</div>

## Installation
```go get github.com/antonsegerkvist/regression```
or use your favorite dependency management tool.

## API Reference
- [Linear](#regressionLinear)
- [Polynomial](#regressionPolynomial)
- [Exponential](#regressionExponential)

<a name="regressionLinear" href="#regressionLinear">#</a> regression.<b>NewLinearRegression2D32</b>() ·
[Source](https://github.com/antonsegerkvist/regression/blob/master/regression.go "Source"),
[Example](https://github.com/antonsegerkvist/regression/blob/master/examples "Example")
<br>
<a href="#regressionLinear">#</a> regression.<b>NewLinearRegression2D64</b>() ·
[Source](https://github.com/antonsegerkvist/regression/blob/master/regression.go "Source"),
[Example](https://github.com/antonsegerkvist/regression/blob/master/examples "Example")

Creates a new linear regression structure with a zero coefficient vector.

![Linear regression](assets/linear-regression.png "Linear regression")


<a name="regressionPolynomial" href="#regressionPolynomial">#</a> regression.<b>NewPolynomialRegression2D64</b>() ·
[Source](https://github.com/antonsegerkvist/regression/blob/master/regression.go "Source"),
[Example](https://github.com/antonsegerkvist/regression/blob/master/examples "Example")
<br>
<a href="#regressionPolynomial">#</a> regression.<b>NewPolynomialRegression2D64</b>() ·
[Source](https://github.com/antonsegerkvist/regression/blob/master/regression.go "Source"),
[Example](https://github.com/antonsegerkvist/regression/blob/master/examples "Example")

Creates a new polynomial regression structure with a zero coefficient vector.

![Polynomial regression](assets/polynomial-regression.png "Polynomial regression")


<a name="regressionExponential" href="#regressionExponential">#</a> regression.<b>NewExponentialRegression2D32</b>() ·
[Source](https://github.com/antonsegerkvist/regression/blob/master/regression.go "Source"),
[Example](https://github.com/antonsegerkvist/regression/blob/master/examples "Example")
<br>
<a href="#regressionExponential">#</a> regression.<b>NewExponentialRegression2D64</b>() ·
[Source](https://github.com/antonsegerkvist/regression/blob/master/regression.go "Source"),
[Example](https://github.com/antonsegerkvist/regression/blob/master/examples "Example")

Creates a new exponential regression structure with a zero coefficient vector.

![Exponential regression](assets/exponential-regression.png "Exponential regression")