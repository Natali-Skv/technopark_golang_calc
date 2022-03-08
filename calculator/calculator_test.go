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
