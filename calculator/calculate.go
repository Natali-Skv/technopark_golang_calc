package calculator

import (
	"fmt"
	"strconv"
)

var precedenceTable = map[tokenType]int{
	multiplication: 3,
	division:       3,
	addition:       2,
	subtraction:    2,
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
	tokens, err := scanExpression([]rune(expression))
	if err != nil {
		return 0, err
	}
	calc := newCalculator()
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

	for err == nil {
		err = calc.performOperation()
	}

	if len(calc.operandStack.operands) != 1 || len(calc.operatorStack.operators) > 0 {
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
	default:
		err = fmt.Errorf("no such operation")
	}
	return err
}
