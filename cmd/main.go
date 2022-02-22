package main

import (
	"calculator"
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
)

func executor(s string) {
	s = strings.TrimSpace(s)
	if s == "" {
		return
	}
	if s == "exit" || s == "quit" {
		os.Exit(0)
	}
	val, err := calculator.Calculate(s)
	if err == nil {
		fmt.Printf("%v\n", val)
	} else {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}
func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "e", Description: ""},
		{Text: "pi", Description: ""},
		{Text: "phi", Description: ""},
		{Text: "sqrt2", Description: ""},
		{Text: "sqrte", Description: ""},
		{Text: "sqrtpi", Description: ""},
		{Text: "sqrtphi", Description: ""},
		{Text: "ln2", Description: ""},
		{Text: "log2e", Description: ""},
		{Text: "ln10", Description: ""},
		{Text: "log10e", Description: ""},
		{Text: "abs", Description: "the absolute value of x"},
		{Text: "acos", Description: "the arccosine, in radians, of x"},
		{Text: "acosh", Description: "the inverse hyperbolic cosine of x"},
		{Text: "asin", Description: "the arcsine, in radians, of x"},
		{Text: "asinh", Description: "the inverse hyperbolic sine of x"},
		{Text: "atan", Description: "the arctangent, in radians, of x"},
		{Text: "atan2", Description: "the arc tangent of y/x, using the signs of the two to determine the quadrant of the return value"},
		{Text: "atanh", Description: "the inverse hyperbolic tangent of x"},
		{Text: "cbrt", Description: "the cube root of x"},
		{Text: "ceil", Description: "the cube root of x"},
		{Text: "copysign", Description: "the cube root of x"},
		{Text: "cos", Description: "the cube root of x"},
		{Text: "cosh", Description: "the hyperbolic cosine of x"},
		{Text: "dim", Description: "the maximum of x-y or 0"},
		{Text: "erf", Description: "the error function of x"},
		{Text: "erfc", Description: "the complementary error function of x"},
		{Text: "erfcinv", Description: "the inverse of Erfc(x)"},
		{Text: "erfinv", Description: "the inverse error function of x"},
		{Text: "exp", Description: "e**x, the base-e exponential of x"},
		{Text: "exp2", Description: "2**x, the base-2 exponential of x"},
		{Text: "expm1", Description: "e**x - 1, the base-e exponential of x minus 1"},
		{Text: "floor", Description: "the greatest integer value less than or equal to x"},
		{Text: "gamma", Description: "the Gamma function of x"},
		{Text: "hypot", Description: "Sqrt(p*p + q*q)"},
		{Text: "j0", Description: "the order-zero Bessel function of the first kind"},
		{Text: "j1", Description: "the order-one Bessel function of the first kind"},
		{Text: "log", Description: "the natural logarithm of x"},
		{Text: "log10", Description: "the decimal logarithm of x"},
		{Text: "log1p", Description: "the natural logarithm of 1 plus its argument x"},
		{Text: "log2", Description: "the binary logarithm of x"},
		{Text: "logb", Description: "the binary exponent of x"},
		{Text: "max", Description: "the larger of x or y"},
		{Text: "min", Description: "the smaller of x or y"},
		{Text: "mod", Description: "the floating-point remainder of x/y"},
		{Text: "nan", Description: "an IEEE 754 “not-a-number” value"},
		{Text: "nextafter", Description: "the next representable float64 value after x towards y"},
		{Text: "pow", Description: "x**y, the base-x exponential of y"},
		{Text: "remainder", Description: "the IEEE 754 floating-point remainder of x/y"},
		{Text: "round", Description: "the nearest integer, rounding half away from zero"},
		{Text: "roundtoeven", Description: "the nearest integer, rounding ties to even"},
		{Text: "sin", Description: "the sine of the radian argument x"},
		{Text: "sinh", Description: "the hyperbolic sine of x"},
		{Text: "sqrt", Description: "the square root of x"},
		{Text: "tan", Description: "the tangent of the radian argument x"},
		{Text: "tanh", Description: "the hyperbolic tangent of x"},
		{Text: "trunc", Description: "the integer value of x"},
		{Text: "y0", Description: "the order-zero Bessel function of the second kind"},
		{Text: "y1", Description: "the order-one Bessel function of the second kind"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix("calculator>> "),
	)
	p.Run()
}
