package algo

import (
	"fmt"
	"unicode"
)

func isDigit(c rune) bool {
	isDigit := unicode.IsDigit(c)
	if isDigit {
		return true
	}
	return false
}

func isLetter(c rune) bool {
	switch c {
	case '*':
		return true
	case '/':
		return true
	case '-':
		return true
	case '+':
		return true
	case '(':
		return true
	case ')':
		return true
	//case ' ':
	//	return true
	default:
		return false
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

	for _, value := range polish {
		if isDigit(value) {
			outputStr = append(outputStr, value)
			outputStr = append(outputStr, ' ')
		}

		curPrior := prioritet(value)
		if len(stack) == 0 && isLetter(value) {
			stack = append(stack, value)
		}
		less := false
		for _, stcVal := range stack {
			stcPrior := prioritet(stcVal)
			if curPrior > stcPrior {
				less = true
			}
		}
		if less {
			stack = append(stack, ' ')
			stack = append(stack, value)
		}
		fmt.Print(len(stack))

	}

	outputStr = append(outputStr, stack...)
	return string(outputStr)
}
