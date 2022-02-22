package calculator

import (
	"fmt"
	"unicode"
)

type tokenType int

const (
	noSuchToken tokenType = iota
	operand
	lParenthsis
	rParenthsis
	addition
	subtraction
	multiplication
	division
	unaryMinus
	mathFunction
)

type token struct {
	tkType tokenType
	value  []rune
	args   [][]token
}

// TODO: checge S to s
func ScanExpression(strExpr []rune, poz int) (tokens []token, newPoz int, err error) {
	bracketCounter := 0
	for poz < len(strExpr) {
		if bracketCounter < 0 {
			return tokens[:len(tokens)-1], poz - 1, nil
		}
		nextToken, newPoz, err := getNextToken(strExpr, poz)
		if err != nil {
			if bracketCounter == 0 {
				return tokens, poz, nil
			}
			return nil, poz, nil
		}
		switch nextToken.tkType {
		case lParenthsis:
			bracketCounter++
		case rParenthsis:
			bracketCounter--
		}
		if nextToken.tkType != noSuchToken {
			tokens = append(tokens, nextToken)
		}
		poz = newPoz
	}
	if bracketCounter < 0 {
		return tokens[:len(tokens)-1], poz - 1, nil
	}
	if bracketCounter > 0 {
		return nil, poz, fmt.Errorf("uncorrect bracket sequence")
	}
	return tokens, poz, nil
}

func getNextToken(str []rune, poz int) (token, int, error) {
	if str == nil {
		return token{}, poz, fmt.Errorf("nil passed as argument")
	}
	if poz >= len(str) {
		return token{}, poz, fmt.Errorf("end of line reached")
	}
	switch {
	case str[poz] == '+':
		return token{addition, []rune{str[poz]}, nil}, poz + 1, nil
	case str[poz] == '*':
		return token{multiplication, []rune{str[poz]}, nil}, poz + 1, nil
	case str[poz] == '/':
		return token{division, []rune{str[poz]}, nil}, poz + 1, nil
	case str[poz] == '(':
		return token{lParenthsis, []rune{str[poz]}, nil}, poz + 1, nil
	case str[poz] == ')':
		return token{rParenthsis, []rune{str[poz]}, nil}, poz + 1, nil
	case unicode.IsSpace(str[poz]):
		return token{}, poz + 1, nil
	case str[poz] == '-':
		pozLastTokenEnd := poz - 1
		for ; pozLastTokenEnd >= 0 && unicode.IsSpace(str[pozLastTokenEnd]); pozLastTokenEnd-- {
		}
		if pozLastTokenEnd < 0 || str[pozLastTokenEnd] == '(' || str[pozLastTokenEnd] == ',' {
			return token{unaryMinus, []rune{str[poz]}, nil}, poz + 1, nil
		}
		return token{subtraction, []rune{str[poz]}, nil}, poz + 1, nil
	case isAlpha(str[poz]):
		{
			literal, poz, err := scanLiteral(str, poz)
			if err != nil {
				return token{}, poz, err
			}

			if _, exist := mathConstants[string(literal)]; exist {
				return token{tkType: operand, value: literal}, poz, nil
			}

			argsNum, err := argsNum(string(literal))
			if err != nil {
				return token{}, poz, err
			}
			funcToken := token{tkType: mathFunction, value: literal}
			funcToken.args, poz, err = scanFuncArgs(argsNum, str, poz)
			if err != nil {
				return token{}, poz, err
			}
			return funcToken, poz, nil
		}
	case unicode.IsDigit(str[poz]):
		{
			number, poz, err := scanOperand(str, poz)
			if err != nil {
				return token{}, poz, err
			}
			return token{tkType: operand, value: number}, poz, nil
		}
	default:
		return token{}, poz, fmt.Errorf("not valid operand or operator")
	}
}

func scanOperand(str []rune, poz int) (operand []rune, newPoz int, err error) {
	if str == nil {
		return nil, poz, fmt.Errorf("nil passed as argument")
	}
	if poz >= len(str) {
		return nil, poz, fmt.Errorf("end of line reached")
	}
	for ; poz < len(str) && (unicode.IsDigit(str[poz]) || str[poz] == '.'); poz++ {
		operand = append(operand, str[poz])
	}
	return operand, poz, nil
}

func scanLiteral(str []rune, poz int) (literal []rune, newPoz int, err error) {
	if str == nil {
		return nil, poz, fmt.Errorf("nil passed as argument")
	}
	if poz >= len(str) {
		return nil, poz, fmt.Errorf("end of line reached")
	}
	if !isAlpha(str[poz]) {
		return nil, poz, fmt.Errorf("it is not math literal")
	}
	pozIter := poz
	for ; pozIter < len(str) && (isAlNum(str[pozIter])); pozIter++ {
		literal = append(literal, unicode.ToLower(str[pozIter]))
	}
	return literal, pozIter, nil
}

func isAlpha(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func isAlNum(char rune) bool {
	return isAlpha(char) || (char >= '0' && char <= '9')
}

func argsNum(funcName string) (int, error) {
	mathFunc, exits := mathFunctions[funcName]
	if !exits {
		return 0, fmt.Errorf("unknown function %s", funcName)
	}
	switch mathFunc.(type) {
	case func() float64:
		return 0, nil
	case func(float64) float64:
		return 1, nil
	case func(float64, float64) float64:
		return 2, nil
	case func(float64, float64, float64) float64:
		return 3, nil
	default:
		return 0, fmt.Errorf("invalid function %s", funcName)
	}
}

func scanFuncArgs(argsNum int, str []rune, poz int) (args [][]token, newPoz int, err error) {
	if str == nil {
		return nil, poz, fmt.Errorf("nil passed as argument")
	}
	if poz >= len(str) {
		return nil, poz, fmt.Errorf("end of line reached")
	}
	if str[poz] != '(' {
		return nil, poz, fmt.Errorf("function arguments must be enclosed in parentheses")
	}
	newPoz = poz + 1
	var currArgument []token
	for ; argsNum > 0 && newPoz < len(str); argsNum-- {
		currArgument, newPoz, err = ScanExpression(str, newPoz)
		if err != nil {
			return nil, newPoz, fmt.Errorf("unncorrect arguments")
		}
		if argsNum > 1 {
			if newPoz < len(str) && str[newPoz] == ',' {
				newPoz++
			} else {
				return nil, newPoz, fmt.Errorf("unncorrect arguments")
			}
		}
		args = append(args, currArgument)
	}
	if newPoz >= len(str) {
		return nil, newPoz, fmt.Errorf("unncorrect arguments")
	}
	nextToken, newPoz, err := getNextToken(str, newPoz)
	if err != nil || nextToken.tkType != rParenthsis {
		return nil, newPoz, fmt.Errorf("unncorrect arguments")
	}
	return args, newPoz, nil
}
