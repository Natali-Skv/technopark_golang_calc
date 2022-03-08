package calculator

import (
	"fmt"
)

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
