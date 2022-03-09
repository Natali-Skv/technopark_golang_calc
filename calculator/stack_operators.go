package calculator

import (
	"fmt"
)

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
