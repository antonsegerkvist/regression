
<div align="center">
<h1>regression-go</h1>
<br/>
<br/>
<p>
regression-go is a Go package containing a collection of linear least-squares fitting methods for simple data analysis.
</p>
</div>

## Installation
```go get github.com/antonsegerkvist/regression-go```
or use your favorite dependency management tool.

## API Reference
- [Linear](#regressionLinear)
- [Exponential](#regressionExponential)

<a name="regressionLinear" href="#regressionLinear">#</a> regression.<b>NewLinearRegression2D32</b>() ·
[Source](https://github.com/antonsegerkvist/regression-go/blob/master/regression.go "Source"),
[Example](https://github.com/antonsegerkvist/regression-go/blob/master/examples "Example")
<br>
<a name="regressionLinear" href="#regressionLinear">#</a> regression.<b>NewLinearRegression2D64</b>() ·
[Source](https://github.com/antonsegerkvist/regression-go/blob/master/regression.go "Source"),
[Example](https://github.com/antonsegerkvist/regression-go/blob/master/examples "Example")

Creates a new linear regression structure with a zero coefficient vector.

![Linear regression](assets/linear-regression.png "Linear regression")

<a name="regressionLinear" href="#regressionLinear">#</a> regression.<b>NewExponentialRegression2D32</b>() ·
[Source](https://github.com/antonsegerkvist/regression-go/blob/master/regression.go "Source"),
[Example](https://github.com/antonsegerkvist/regression-go/blob/master/examples "Example")
<br>
<a name="regressionLinear" href="#regressionLinear">#</a> regression.<b>NewExponentialRegression2D64</b>() ·
[Source](https://github.com/antonsegerkvist/regression-go/blob/master/regression.go "Source"),
[Example](https://github.com/antonsegerkvist/regression-go/blob/master/examples "Example")

Creates a new exponential regression structure with a zero coefficient vector.

![Exponential regression](assets/exponential-regression.png "Exponential regression")