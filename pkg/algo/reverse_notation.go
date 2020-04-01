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

func reverseStack(s []rune) []rune {
	var result []rune
	if len(s) == 1 {
		return s
	}
	for i := len(s) - 1; i >= 0; i-- {
		fmt.Print(s[i])
		result = append(result, s[i])
	}
	return result
}

func isSpace(c rune) bool {
	if c == ' ' {
		return true
	}
	return false
}

func checkStack(stack []rune) bool {
	if len(stack) == 0 {
		return true
	}
	return false
}

func isOpenBrackets(c rune) bool {
	switch c {
	case '(':
		return true
	default:
		return false
	}
}

func isCloseBrackets(c rune) bool {
	switch c {
	case ')':
		return true
	default:
		return false
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
			} else {
				less = false
				break
			}
		}
		if less {
			stack = append(stack, ' ')
			stack = append(stack, value)
		}
		fmt.Print(len(stack))
	}

	outputStr = append(outputStr, reverseStack(stack)...)
	return string(outputStr)
}
