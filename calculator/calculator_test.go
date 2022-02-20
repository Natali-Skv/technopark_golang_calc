package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculation(t *testing.T) {
	tests := []struct {
		in          string
		expectedOut float64
	}{
		{"-2.5", -2.5},
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

// TODO: not-valid data

// func TestNan(t *testing.T) {
// 	val, err := Calculate("nan()")
// 	if assert.NoError(t, err) {
// 		assert.True(t, math.IsNaN(val))
// 	}
// }
