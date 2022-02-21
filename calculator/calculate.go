package calculator

import (
	"fmt"
	"math"
	"strconv"
)

var precedenceTable = map[tokenType]int{
	unaryMinus:     3,
	multiplication: 2,
	division:       2,
	addition:       1,
	subtraction:    1,
}

var mathFunctions = map[string]interface{}{
	"abs":         math.Abs,
	"acos":        math.Acos,
	"acosh":       math.Acosh,
	"asin":        math.Asin,
	"asinh":       math.Asinh,
	"atan":        math.Atan,
	"atan2":       math.Atan2,
	"atanh":       math.Atanh,
	"cbrt":        math.Cbrt,
	"ceil":        math.Ceil,
	"copysign":    math.Copysign,
	"cos":         math.Cos,
	"cosh":        math.Cosh,
	"dim":         math.Dim,
	"erf":         math.Erf,
	"erfc":        math.Erfc,
	"erfcinv":     math.Erfcinv, // Go 1.10+
	"erfinv":      math.Erfinv,  // Go 1.10+
	"exp":         math.Exp,
	"exp2":        math.Exp2,
	"expm1":       math.Expm1,
	"floor":       math.Floor,
	"gamma":       math.Gamma,
	"hypot":       math.Hypot,
	"j0":          math.J0,
	"j1":          math.J1,
	"log":         math.Log,
	"log10":       math.Log10,
	"log1p":       math.Log1p,
	"log2":        math.Log2,
	"logb":        math.Logb,
	"max":         math.Max,
	"min":         math.Min,
	"mod":         math.Mod,
	"nan":         math.NaN,
	"nextafter":   math.Nextafter,
	"pow":         math.Pow,
	"remainder":   math.Remainder,
	"round":       math.Round,       // Go 1.10+
	"roundtoeven": math.RoundToEven, // Go 1.10+
	"sin":         math.Sin,
	"sinh":        math.Sinh,
	"sqrt":        math.Sqrt,
	"tan":         math.Tan,
	"tanh":        math.Tanh,
	"trunc":       math.Trunc,
	"y0":          math.Y0,
	"y1":          math.Y1,
}

type stackOperators struct {
	operators []token
}

func (s *stackOperators) pop() (token, error) {
	if len(s.operators) <= 0 {
		return token{}, fmt.Errorf("stack is empty")
	}
	top := s.operators[len(s.operators)-1]
	s.operators = s.operators[:len(s.operators)-1]
	return top, nil
}

func (s *stackOperators) top() (token, error) {
	if len(s.operators) <= 0 {
		return token{}, fmt.Errorf("stack is empty")
	}
	return s.operators[len(s.operators)-1], nil
}

func (s *stackOperators) push(newOperand token) {
	s.operators = append(s.operators, newOperand)
}

type stackOperands struct {
	operands []float64
}

func (s *stackOperands) pop(count int) ([]float64, error) {
	if len(s.operands)-count < 0 {
		return nil, fmt.Errorf("stack is empty")
	}
	topCount := s.operands[len(s.operands)-count:]
	s.operands = s.operands[:len(s.operands)-count]
	return topCount, nil
}

func (s *stackOperands) push(newOperand float64) {
	s.operands = append(s.operands, newOperand)
}

type calcuator struct {
	operatorStack stackOperators
	operandStack  stackOperands
}

func newCalculator() *calcuator {
	return &calcuator{
		operatorStack: stackOperators{operators: make([]token, 0, 10)},
		operandStack:  stackOperands{operands: make([]float64, 0, 10)},
	}
}

func Calculate(expression string) (float64, error) {
	tokens, indEndExpr, err := ScanExpression([]rune(expression), 0)
	if err != nil || indEndExpr < len(expression) {
		return 0, fmt.Errorf("uncorrect expression")
	}
	calc := newCalculator()
	return calc.calculate(tokens)
}

func (calc *calcuator) calculate(tokens []token) (float64, error) {
	var err error
	for _, currToken := range tokens {
		switch currToken.tkType {
		case operand:
			{
				operandNumber, err := strconv.ParseFloat(string(currToken.value), 64)
				if err != nil {
					return 0, err
				}
				calc.operandStack.push(operandNumber)
			}
		case rParenthsis:
			{
				err = calc.popUtilLeftParenthsis()
				if err != nil {
					return 0, err
				}
			}
		case lParenthsis:
			{
				calc.operatorStack.push(currToken)
			}
		case mathFunction:
			{
				var args []float64
				for _, arg := range currToken.args {
					calcArg := newCalculator()
					argRes, err := calcArg.calculate(arg)
					if err != nil {
						return 0, err
					}
					args = append(args, argRes)
				}
				var result float64
				result, err = callMathFunc(string(currToken.value), args)
				calc.operandStack.push(result)
			}
		default:
			{
				currOperator := currToken

				for topOperator, err := calc.operatorStack.top(); err == nil && precedenceTable[topOperator.tkType] >= precedenceTable[currOperator.tkType]; topOperator, err = calc.operatorStack.top() {
					calc.performOperation()
				}
				calc.operatorStack.push(currToken)
			}
		}
	}

	for len(calc.operatorStack.operators) > 0 && err == nil {
		err = calc.performOperation()
	}

	if len(calc.operandStack.operands) != 1 || len(calc.operatorStack.operators) > 0 || err != nil {
		return 0, fmt.Errorf("expression is not valid")
	}
	result, err := calc.operandStack.pop(1)
	return result[0], err
}

func (calc *calcuator) popUtilLeftParenthsis() error {
	operator, err := calc.operatorStack.top()
	for ; err == nil && operator.tkType != lParenthsis; operator, err = calc.operatorStack.top() {
		err = calc.performOperation()
		if err != nil {
			return fmt.Errorf("parentheses incorrectly placed")
		}
	}
	if err != nil {
		return fmt.Errorf("parentheses incorrectly placed")
	}
	_, err = calc.operatorStack.pop()
	return err
}

func (calc *calcuator) performOperation() error {
	operator, err := calc.operatorStack.pop()
	if err != nil {
		return err
	}
	switch operator.tkType {
	case multiplication:
		{
			operands, err := calc.operandStack.pop(2)
			if err != nil {
				return err
			}
			lOperand, rOperand := operands[0], operands[1]

			calc.operandStack.push(lOperand * rOperand)
		}
	case division:
		{
			operands, err := calc.operandStack.pop(2)
			if err != nil {
				return err
			}
			lOperand, rOperand := operands[0], operands[1]

			calc.operandStack.push(lOperand / rOperand)
		}
	case addition:
		{
			operands, err := calc.operandStack.pop(2)
			if err != nil {
				return err
			}
			lOperand, rOperand := operands[0], operands[1]
			calc.operandStack.push(lOperand + rOperand)
		}
	case subtraction:
		{
			operands, err := calc.operandStack.pop(2)
			if err != nil {
				return err
			}
			lOperand, rOperand := operands[0], operands[1]
			calc.operandStack.push(lOperand - rOperand)
		}
	case unaryMinus:
		{
			operands, err := calc.operandStack.pop(1)
			if err != nil {
				return err
			}
			operand := operands[0]
			calc.operandStack.push(-operand)
		}
	default:
		err = fmt.Errorf("no such operation")
	}
	return err
}

func callMathFunc(funcName string, args []float64) (float64, error) {
	mathFunc, exists := mathFunctions[funcName]
	if !exists {
		return 0, fmt.Errorf("unknown function %s", funcName)
	}
	switch mathFunc := mathFunc.(type) {
	case func() float64:
		return mathFunc(), nil
	case func(float64) float64:
		return mathFunc(args[0]), nil
	case func(float64, float64) float64:
		return mathFunc(args[0], args[1]), nil
	case func(float64, float64, float64) float64:
		return mathFunc(args[0], args[1], args[2]), nil
	default:
		return 0, fmt.Errorf("invalid function %s", funcName)
	}
}
