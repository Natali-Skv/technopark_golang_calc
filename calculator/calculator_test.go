package calculator

import (
	"fmt"
	"math"
	"math/rand"
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

func TestStackOperandsPop(t *testing.T) {
	testCases := []struct {
		in []float64
	}{
		{[]float64{123, 33, 10.1010, 11, 111. - 10, 11, 332, 10.3, 40.09, 500, -6, 97, 88, 89}},
		{[]float64{-10, 11, 332, 10.3, 40.09, 500, -6, 97, 88, 89}},
		{[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]float64{-10, -10, -10, -10}},
		{[]float64{0}},
		{[]float64{}},
	}

	for _, testCase := range testCases {

		for i := 0; i < len(testCase.in); i++ {
			t.Run(fmt.Sprint(testCase.in), func(t *testing.T) {
				testStack := stackOperands{testCase.in}
				popLen := rand.Intn(len(testCase.in))
				result, err := testStack.pop(popLen)
				if assert.NoError(t, err) {
					assert.Equal(t, testCase.in[len(testCase.in)-popLen:], result)
				}
			})
		}
	}
}

func TestStackOperandsPush(t *testing.T) {
	for i := 0; i < 10; i++ {
		testStack := stackOperands{}
		numToPush := rand.Intn(10) + 1
		expectStack := make([]float64, 0, numToPush)

		for j := 0; j < numToPush; j++ {
			valToPush := rand.Float64()*100 - 100
			testStack.push(valToPush)
			expectStack = append(expectStack, valToPush)
		}
		t.Run(fmt.Sprint(expectStack), func(t *testing.T) {
			result, err := testStack.pop(numToPush)
			if assert.NoError(t, err) {
				assert.Equal(t, expectStack, result)
			}
		})
	}
}

func TestStackOperandsInvalid(t *testing.T) {
	testStack := stackOperands{}

	result, err := testStack.pop(1)
	assert.Nil(t, result, "Return slice value is not nil, but pop from emty stack executed")
	assert.NotNil(t, err, "Error return value is nil, but expected some error, pop from empty stack")
}

func TestScanExpresion(t *testing.T) {
	type structIn struct {
		expr []rune
		poz  int
	}
	type expectStruct struct {
		tokens []token
		newPoz int
	}

	testCases := []struct {
		in           structIn
		expectTokens expectStruct
	}{
		{
			in:           structIn{expr: []rune("1"), poz: 0},
			expectTokens: expectStruct{tokens: []token{{operand, []rune("1"), [][]token(nil)}}, newPoz: 1},
		},
		{
			in:           structIn{expr: []rune("1+1"), poz: 0},
			expectTokens: expectStruct{tokens: []token{{operand, []rune("1"), [][]token(nil)}, {addition, []rune("+"), [][]token(nil)}, {operand, []rune("1"), [][]token(nil)}}, newPoz: 3},
		},
		{
			in:           structIn{expr: []rune("(12+13)*2"), poz: 0},
			expectTokens: expectStruct{tokens: []token{{lParenthsis, []rune("("), [][]token(nil)}, {operand, []rune("12"), [][]token(nil)}, {addition, []rune("+"), [][]token(nil)}, {operand, []rune("13"), [][]token(nil)}, {rParenthsis, []rune(")"), [][]token(nil)}, {multiplication, []rune("*"), [][]token(nil)}, {operand, []rune("2"), [][]token(nil)}}, newPoz: 9},
		},
		{
			in:           structIn{expr: []rune("(12+13)*2/1.5"), poz: 0},
			expectTokens: expectStruct{tokens: []token{{lParenthsis, []rune("("), [][]token(nil)}, {operand, []rune("12"), [][]token(nil)}, {addition, []rune("+"), [][]token(nil)}, {operand, []rune("13"), [][]token(nil)}, {rParenthsis, []rune(")"), [][]token(nil)}, {multiplication, []rune("*"), [][]token(nil)}, {operand, []rune("2"), [][]token(nil)}, {division, []rune("/"), [][]token(nil)}, {operand, []rune("1.5"), [][]token(nil)}}, newPoz: 13},
		},
		{
			in:           structIn{expr: []rune("calculate:(12+13)*2/1.5"), poz: 10},
			expectTokens: expectStruct{tokens: []token{{lParenthsis, []rune("("), [][]token(nil)}, {operand, []rune("12"), [][]token(nil)}, {addition, []rune("+"), [][]token(nil)}, {operand, []rune("13"), [][]token(nil)}, {rParenthsis, []rune(")"), [][]token(nil)}, {multiplication, []rune("*"), [][]token(nil)}, {operand, []rune("2"), [][]token(nil)}, {division, []rune("/"), [][]token(nil)}, {operand, []rune("1.5"), [][]token(nil)}}, newPoz: 23},
		},
		{
			in:           structIn{expr: []rune("calculate: ( -12 + 13 ) * 2   / 1.5 please"), poz: 11},
			expectTokens: expectStruct{tokens: []token{{lParenthsis, []rune("("), [][]token(nil)}, {unaryMinus, []rune("-"), [][]token(nil)}, {operand, []rune("12"), [][]token(nil)}, {addition, []rune("+"), [][]token(nil)}, {operand, []rune("13"), [][]token(nil)}, {rParenthsis, []rune(")"), [][]token(nil)}, {multiplication, []rune("*"), [][]token(nil)}, {operand, []rune("2"), [][]token(nil)}, {division, []rune("/"), [][]token(nil)}, {operand, []rune("1.5"), [][]token(nil)}}, newPoz: 36},
		},

		{
			in:           structIn{expr: []rune("pi"), poz: 0},
			expectTokens: expectStruct{tokens: []token{{operand, []rune("pi"), [][]token(nil)}}, newPoz: 2},
		},
		{
			in:           structIn{expr: []rune("cos(1)"), poz: 0},
			expectTokens: expectStruct{tokens: []token{{mathFunction, []rune("cos"), [][]token{{{operand, []rune("1"), [][]token(nil)}}}}}, newPoz: 6},
		},

		{
			in:           structIn{expr: []rune("cos((12+13)*2/1.5)"), poz: 0},
			expectTokens: expectStruct{tokens: []token{{mathFunction, []rune("cos"), [][]token{{{lParenthsis, []rune("("), [][]token(nil)}, {operand, []rune("12"), [][]token(nil)}, {addition, []rune("+"), [][]token(nil)}, {operand, []rune("13"), [][]token(nil)}, {rParenthsis, []rune(")"), [][]token(nil)}, {multiplication, []rune("*"), [][]token(nil)}, {operand, []rune("2"), [][]token(nil)}, {division, []rune("/"), [][]token(nil)}, {operand, []rune("1.5"), [][]token(nil)}}}}}, newPoz: 18},
		},
		{
			in:           structIn{expr: []rune("sin(cos((12+13)*2/1.5))"), poz: 0},
			expectTokens: expectStruct{tokens: []token{{mathFunction, []rune("sin"), [][]token{{{mathFunction, []rune("cos"), [][]token{{{lParenthsis, []rune("("), [][]token(nil)}, {operand, []rune("12"), [][]token(nil)}, {addition, []rune("+"), [][]token(nil)}, {operand, []rune("13"), [][]token(nil)}, {rParenthsis, []rune(")"), [][]token(nil)}, {multiplication, []rune("*"), [][]token(nil)}, {operand, []rune("2"), [][]token(nil)}, {division, []rune("/"), [][]token(nil)}, {operand, []rune("1.5"), [][]token(nil)}}}}}}}}, newPoz: 23},
		},

		{
			in:           structIn{expr: []rune("max(1,2)"), poz: 0},
			expectTokens: expectStruct{tokens: []token{{mathFunction, []rune("max"), [][]token{{{operand, []rune("1"), [][]token(nil)}}, {{operand, []rune("2"), [][]token(nil)}}}}}, newPoz: 8},
		},

		{
			in:           structIn{expr: []rune("max(-1.5,2)"), poz: 0},
			expectTokens: expectStruct{tokens: []token{{mathFunction, []rune("max"), [][]token{{{unaryMinus, []rune("-"), [][]token(nil)}, {operand, []rune("1.5"), [][]token(nil)}}, {{operand, []rune("2"), [][]token(nil)}}}}}, newPoz: 11},
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprint(testCase.in), func(t *testing.T) {
			resultTokens, resultNewPoz, err := scanExpression(testCase.in.expr, testCase.in.poz)
			if assert.NoError(t, err) {
				assert.Equal(t, testCase.expectTokens.tokens, resultTokens)
				assert.Equal(t, testCase.expectTokens.newPoz, resultNewPoz)
			}
		})
	}
}

func TestScanExpresionInvalid(t *testing.T) {

	testCases := []struct {
		expr   []rune
		poz    int
		newPoz int
	}{
		{
			expr:   []rune(""),
			poz:    0,
			newPoz: 0,
		},
		{
			expr:   []rune(""),
			poz:    10,
			newPoz: 10,
		},
		{
			expr:   []rune("xx(1)"),
			poz:    0,
			newPoz: 0,
		},
		{
			expr:   []rune("^10"),
			poz:    0,
			newPoz: 0,
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprint(testCase), func(t *testing.T) {
			resultTokens, resultNewPoz, err := scanExpression(testCase.expr, testCase.poz)
			assert.NoError(t, err)
			assert.Nil(t, resultTokens)
			assert.Equal(t, testCase.newPoz, resultNewPoz)
		})
	}
}

func TestCalculete(t *testing.T) {
	testCases := []struct {
		in           []token
		expectResult float64
	}{
		{
			in:           []token{{operand, []rune("1"), [][]token(nil)}},
			expectResult: 1,
		},
		{
			in:           []token{{operand, []rune("1"), [][]token(nil)}, {addition, []rune("+"), [][]token(nil)}, {operand, []rune("1"), [][]token(nil)}},
			expectResult: 2,
		},
		{
			in:           []token{{lParenthsis, []rune("("), [][]token(nil)}, {operand, []rune("12"), [][]token(nil)}, {addition, []rune("+"), [][]token(nil)}, {operand, []rune("13"), [][]token(nil)}, {rParenthsis, []rune(")"), [][]token(nil)}, {multiplication, []rune("*"), [][]token(nil)}, {operand, []rune("2"), [][]token(nil)}},
			expectResult: 50,
		},
		{
			in:           []token{{lParenthsis, []rune("("), [][]token(nil)}, {operand, []rune("12"), [][]token(nil)}, {addition, []rune("+"), [][]token(nil)}, {operand, []rune("13"), [][]token(nil)}, {rParenthsis, []rune(")"), [][]token(nil)}, {multiplication, []rune("*"), [][]token(nil)}, {operand, []rune("2"), [][]token(nil)}, {division, []rune("/"), [][]token(nil)}, {operand, []rune("1.5"), [][]token(nil)}},
			expectResult: 33.333333333333,
		},
		{
			in:           []token{{lParenthsis, []rune("("), [][]token(nil)}, {unaryMinus, []rune("-"), [][]token(nil)}, {operand, []rune("12"), [][]token(nil)}, {addition, []rune("+"), [][]token(nil)}, {operand, []rune("13"), [][]token(nil)}, {rParenthsis, []rune(")"), [][]token(nil)}, {multiplication, []rune("*"), [][]token(nil)}, {operand, []rune("2"), [][]token(nil)}, {division, []rune("/"), [][]token(nil)}, {operand, []rune("1.5"), [][]token(nil)}},
			expectResult: 1.3333333333333,
		},

		{
			in:           []token{{operand, []rune("pi"), [][]token(nil)}},
			expectResult: math.Pi,
		},

		{
			in:           []token{{mathFunction, []rune("max"), [][]token{{{operand, []rune("1"), [][]token(nil)}}, {{operand, []rune("2"), [][]token(nil)}}}}},
			expectResult: 2,
		},

		{
			in:           []token{{mathFunction, []rune("max"), [][]token{{{unaryMinus, []rune("-"), [][]token(nil)}, {operand, []rune("1.5"), [][]token(nil)}}, {{operand, []rune("2"), [][]token(nil)}}}}},
			expectResult: 2,
		},
	}

	calc := newCalculator()
	for _, testCase := range testCases {
		t.Run(fmt.Sprint(testCase.in), func(t *testing.T) {
			result, err := calc.calculate(testCase.in)
			if assert.NoError(t, err) {
				assert.InDelta(t, testCase.expectResult, result, 0.000001)
			}
		})
	}
}

func TestCalculeteInvalid(t *testing.T) {
	testCases := []struct {
		in      []token
		errType string
	}{
		{
			in:      []token{{operand, []rune("x"), [][]token(nil)}},
			errType: "not operand",
		},
		{
			in:      []token{{operand, []rune("a1"), [][]token(nil)}, {addition, []rune("+"), [][]token(nil)}, {operand, []rune("1"), [][]token(nil)}},
			errType: "not operand",
		},
		{
			in:      []token{{lParenthsis, []rune("("), [][]token(nil)}, {lParenthsis, []rune("("), [][]token(nil)}, {operand, []rune("12"), [][]token(nil)}, {addition, []rune("+"), [][]token(nil)}, {operand, []rune("13"), [][]token(nil)}, {rParenthsis, []rune(")"), [][]token(nil)}, {multiplication, []rune("*"), [][]token(nil)}, {operand, []rune("2"), [][]token(nil)}},
			errType: "wrong parenthsis sequence",
		},
		{
			in:      []token{{rParenthsis, []rune(")"), [][]token(nil)}, {operand, []rune("12"), [][]token(nil)}, {addition, []rune("+"), [][]token(nil)}, {operand, []rune("13"), [][]token(nil)}, {rParenthsis, []rune(")"), [][]token(nil)}, {multiplication, []rune("*"), [][]token(nil)}, {operand, []rune("2"), [][]token(nil)}, {division, []rune("/"), [][]token(nil)}, {operand, []rune("1.5"), [][]token(nil)}},
			errType: "wrong parenthsis sequence",
		},
		{
			in:      []token{{lParenthsis, []rune("("), [][]token(nil)}, {unaryMinus, []rune("-"), [][]token(nil)}, {unaryMinus, []rune("-"), [][]token(nil)}, {operand, []rune("12"), [][]token(nil)}, {addition, []rune("+"), [][]token(nil)}, {operand, []rune("13"), [][]token(nil)}, {rParenthsis, []rune(")"), [][]token(nil)}, {multiplication, []rune("*"), [][]token(nil)}, {operand, []rune("2"), [][]token(nil)}, {division, []rune("/"), [][]token(nil)}, {operand, []rune("1.5"), [][]token(nil)}},
			errType: "two minus",
		},

		{
			in:      []token{{mathFunction, []rune("no"), [][]token{{{operand, []rune("1"), [][]token(nil)}}, {{operand, []rune("2"), [][]token(nil)}}}}},
			errType: "no such function",
		},

		{
			in:      []token{{mathFunction, []rune("max"), [][]token{{{unaryMinus, []rune("-"), [][]token(nil)}, {operand, []rune("1.5"), [][]token(nil)}}, {{operand, []rune("2"), [][]token(nil)}}, {{operand, []rune("2"), [][]token(nil)}}}}},
			errType: "wrang argument count",
		},
	}

	calc := newCalculator()
	for _, testCase := range testCases {
		t.Run(fmt.Sprint(testCase.errType), func(t *testing.T) {
			result, err := calc.calculate(testCase.in)
			assert.Error(t, err)
			assert.Equal(t, 0.0, result)
		})
	}
}
