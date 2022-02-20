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
}

func scanExpression(strExpr []rune) ([]token, error) {
	tokens := make([]token, 0, len(strExpr))
	for poz := 0; poz < len(strExpr); {
		nextToken, newPoz, err := getNextToken(strExpr, poz)
		if err != nil {
			return nil, err
		}
		if nextToken.tkType != noSuchToken {
			tokens = append(tokens, nextToken)
		}
		poz = newPoz
	}
	return tokens, nil
}

//  TODO add functions sin ect

func getNextToken(str []rune, poz int) (token, int, error) {
	switch {
	case str[poz] == '+':
		return token{addition, []rune{str[poz]}}, poz + 1, nil
	case str[poz] == '*':
		return token{multiplication, []rune{str[poz]}}, poz + 1, nil
	case str[poz] == '/':
		return token{division, []rune{str[poz]}}, poz + 1, nil
	case str[poz] == '(':
		return token{lParenthsis, []rune{str[poz]}}, poz + 1, nil
	case str[poz] == ')':
		return token{rParenthsis, []rune{str[poz]}}, poz + 1, nil
	case unicode.IsSpace(str[poz]):
		return token{}, poz + 1, nil
	case str[poz] == '-':
		pozLastTokenEnd := poz - 1
		for ; pozLastTokenEnd >= 0 && unicode.IsSpace(str[pozLastTokenEnd]); pozLastTokenEnd-- {
		}
		if pozLastTokenEnd < 0 || str[pozLastTokenEnd] == '(' {
			return token{unaryMinus, []rune{str[poz]}}, poz + 1, nil
		}
		return token{subtraction, []rune{str[poz]}}, poz + 1, nil
	// case !unicode.IsDigit(str[poz]):
	// return token{}, poz, fmt.Errorf("not valid operand or operator")
	case isAlpha(str[poz]):
		{
			funcName, poz := scanFunc(str, poz)
			return token{tkType: mathFunction, value: funcName}, poz, nil

		}
	case unicode.IsDigit(str[poz]):
		{
			number, poz := scanOperand(str, poz)
			return token{tkType: operand, value: number}, poz, nil
		}
	default:
		return token{}, poz, fmt.Errorf("not valid operand or operator")
	}
}

func scanOperand(str []rune, poz int) (operand []rune, newPoz int) {
	for ; poz < len(str) && (unicode.IsDigit(str[poz]) || str[poz] == '.'); poz++ {
		operand = append(operand, str[poz])
	}
	return operand, poz
}

func scanFunc(str []rune, poz int) (funcName []rune, newPoz int) {
	if !isAlpha(str[poz]) {
		return nil, poz
	}
	for ; poz < len(str) && (isAlNum(str[poz])); poz++ {
		funcName = append(funcName, str[poz])
	}
	return funcName, poz
}

func isAlpha(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func isAlNum(char rune) bool {
	return isAlpha(char) || (char >= '0' && char <= '9')
}