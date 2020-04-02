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
	var curPrior int8

	for _, value := range polish {
		if isDigit(value) {
			outputStr = append(outputStr, value)
			outputStr = append(outputStr, ' ')
		}

		if isLetter(value) {
			curPrior = prioritet(value)
		}

		if checkStack(stack) && isLetter(value) {
			stack = append(stack, value)
		}

		less := false
		if !isDigit(value) {
			for _, stcVal := range stack {
				stcPrior := prioritet(stcVal)
				if (curPrior > stcPrior) && !isSpace(stcVal) {
					less = true
				} else {
					less = false
					break
				}
			}
		}
		if less {
			stack = append(stack, ' ')
			stack = append(stack, value)
		}

		if (len(stack) > 0) && !isDigit(value) {
			fmt.Print(string(stack))
			stack := reverseStack(stack)
			for i, v := range stack {
				if curPrior < prioritet(v) {
					if !isSpace(value) {
						outputStr = append(outputStr, v)
						outputStr = append(outputStr, ' ')
						stack = append(stack[:i], stack[i+1:]...)
						stack = append(stack, value)
					}
				} else {
					break
				}
			}
		}
	}

	outputStr = append(outputStr, reverseStack(stack)...)
	return string(outputStr)
}
