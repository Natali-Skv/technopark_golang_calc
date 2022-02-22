package calculator

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		in          string
		expectedOut float64
	}{
		{"-2.5", -2.5},
		{"-(-2.5)", 2.5},
		{"-(-(-2.5))", -2.5},
		{"1.2+2.3", 1.2 + 2.3},
		{"1.2-2.3", 1.2 - 2.3},
		{"2.5*3", 2.5 * 3},
		{"3.0/2.0", 3.0 / 2.0},
		{"-1.2+2.3", -1.2 + 2.3},
		{"+1.2+2.3", +1.2 + 2.3},
		{"1.2+2.5*2.0", 1.2 + 2.5*2.0},
		{"(1.2+2.5)*2.0", (1.2 + 2.5) * 2.0},
		{" (1.2 +  2.5 ) *   2.0  ", (1.2 + 2.5) * 2.0},
		{" ( -1.2 +  ( -2.5 ) ) *   2.0  ", (-1.2 - 2.5) * 2.0},
		{" (-1.2 +  (-2.5) ) *   2.0  ", (-1.2 - 2.5) * 2.0},
		{"(0+1)*0+1/4+(1*1*1*1*1)*10+100-100", (10.25)},
		{" ( -1.2 +  ( -2.5 ) ) /   0  ", math.Inf(-1)},

		{"e", math.E},
		{"E", math.E},

		{"pi*2.0", math.Pi * 2.0},
		{"Pi*2.0", math.Pi * 2.0},
		{"PI*2.0", math.Pi * 2.0},

		{"sqrt2", math.Sqrt2},
		{"sqrte", math.SqrtE},
		{"sqrtpi", math.SqrtPi},
		{"sqrtphi", math.SqrtPhi},

		{"ln2", math.Ln2},
		{"log2e", math.Log2E},
		{"ln2 * log2e", math.Ln2 * math.Log2E},
		{"ln10", math.Ln10},
		{"log10e", math.Log10E},
		{"ln10 * log10e", math.Ln10 * math.Log10E},

		{"abs(-1.5)", math.Abs(-1.5)},
		{"Abs(-1.5)", math.Abs(1.5)},
		{"ABS(-1.5)", math.Abs(1.5)},
		{"abs( (1.2 +  2.5 ) *   2.0  )", math.Abs((1.2 + 2.5) * 2.0)},

		{"Atan2(1.2,3.4)", math.Atan2(1.2, 3.4)},
		{"Atan2(1.2, 3.4)", math.Atan2(1.2, 3.4)},
		{"Atan2( (1.0 + 0.2) * 0.4, -1.7 * 2)", math.Atan2((1.0+0.2)*0.4, -1.7*2)},
	}
	for _, testCase := range tests {
		t.Run(testCase.in, func(t *testing.T) {
			val, err := Calculate(testCase.in)
			if assert.NoError(t, err) {
				assert.InDelta(t, testCase.expectedOut, val, 0.001)
			}
		})
	}
}

func TestCalculateInvalidArgs(t *testing.T) {
	tests := []struct {
		in          string
		expectedOut float64
	}{
		{"", 0},
		{"(", 0},
		{".1", 0},
		{"3/", 0},
		{"+", 0},
		{")", 0},
		{"(1+1))", 0},
		{"((1*)1)", 0},

		{"pow", 0},
		{"pow(1)", 0},
		{"pow(1,2,3)", 0},
		{"pow(1,2))", 0},
		{"pow((1,2)", 0},
		{"pow((1,2))", 0},
		{"pow(1(,2))", 0},
		{"noSuchFcn(1)", 0},
		{"*pow(1,2)", 0},
	}
	for _, testCase := range tests {
		t.Run(testCase.in, func(t *testing.T) {
			val, err := Calculate(testCase.in)
			if assert.Error(t, err) {
				assert.InDelta(t, testCase.expectedOut, val, 0.001)
			}
		})
	}
}
