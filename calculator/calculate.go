package calculator

import (
	"fmt"
	"strconv"
)

var precedenceTable = map[tokenType]int{
	unaryMinus:     3,
	multiplication: 2,
	division:       2,
	addition:       1,
	subtraction:    1,
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
	tokens, indEndExpr, err := scanExpression([]rune(expression), 0)

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
				var operandNumber float64
				var exists bool
				if operandNumber, exists = mathConstants[string(currToken.value)]; !exists {
					operandNumber, err = strconv.ParseFloat(string(currToken.value), 64)
					if err != nil {
						return 0, err
					}
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
