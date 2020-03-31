package algo

import (
	"unicode"
)

func isDigit(c rune) bool {
	isDigit := unicode.IsDigit(c)
	if isDigit {
		return true
	}
	return false
}

func isLetter(c rune) (bool, string, int8) {
	switch c {
	case '*':
		return true, "mult", 3
	case '/':
		return true, "div", 3
	case '-':
		return true, "minus", 2
	case '+':
		return true, "plus", 2
	case '(':
		return true, "open", 1
	case ')':
		return true, "close", 0
	case ' ':
		return true, "close", -3
	default:
		return false, "", -1
	}
}

func prioritet(c rune) int8 {
	switch c {
	case '*', '/':
		return 3
	case '-', '+':
		return 2
	case '(':
		return 1
	case ')':
		return 0
	default:
		return -1
	}
}

func ReverseNotation(polish string) string {
	var stack []rune
	var outputStr []rune

	for _, val := range polish {
		isLet, _, code := isLetter(val)
		if (isDigit(val)) || (code == -3) {
			outputStr = append(outputStr, val)
		}
		if (isLet) || (code == -3) {
			stack = append(stack, val)
		}
	}
	outputStr = append(outputStr, stack...)

	return string(outputStr)
}
