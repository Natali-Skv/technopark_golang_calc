package calculator

import (
	"fmt"
	"unicode"
)

type tokenType int

const (
	noSuchTokenType tokenType = iota
	operand
	lParenthsis
	rParenthsis
	addition
	subtraction
	multiplication
	division
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
		tokens = append(tokens, nextToken)
		poz = newPoz
	}
	return tokens, nil
}

//  TODO add functions sin ect

func getNextToken(str []rune, poz int) (token, int, error) {
	switch str[poz] {
	case '-':
		return token{subtraction, []rune{str[poz]}}, poz + 1, nil
	case '+':
		return token{addition, []rune{str[poz]}}, poz + 1, nil
	case '*':
		return token{multiplication, []rune{str[poz]}}, poz + 1, nil
	case '/':
		return token{division, []rune{str[poz]}}, poz + 1, nil
	case '(':
		return token{lParenthsis, []rune{str[poz]}}, poz + 1, nil
	case ')':
		return token{rParenthsis, []rune{str[poz]}}, poz + 1, nil
	default:
		if !unicode.IsDigit(str[poz]) {
			return token{}, poz, fmt.Errorf("not valid operand or operator")
		}

		number, poz := scanOperand(str, poz)
		return token{tkType: operand, value: number}, poz, nil
	}
}

func scanOperand(str []rune, poz int) (operand []rune, newPoz int) {
	for ; poz < len(str) && (unicode.IsDigit(str[poz]) || str[poz] == '.'); poz++ {
		operand = append(operand, str[poz])
	}
	return operand, poz
}
