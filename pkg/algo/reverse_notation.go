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
	default:
		return 0
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

func hasOpenBrackets(c []rune) (int, bool) {
	for i, v := range c {
		if isOpenBrackets(v) {
			return i, true
		}
	}
	return -1, false
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
			continue
		}

		if isOpenBrackets(value) {
			stack = append(stack, value)
			continue
		}

		if isCloseBrackets(value) {
			stack = reverseStack(stack)
			idxOpenBrackets, hasBrackets := hasOpenBrackets(stack)
			if hasBrackets {
				outputStr = append(outputStr, stack[:idxOpenBrackets-1]...)
				stack = append(stack[:0], stack[idxOpenBrackets+1:]...)
			}
			stack = reverseStack(stack)
			continue
		}

		if len(stack) != 0 {
			if isOpenBrackets(stack[len(stack)-1]) {
				stack = append(stack, value)
				continue
			}
			if isLetter(value) && (curPrior > prioritet(stack[len(stack)-1])) {
				stack = append(stack, ' ')
				stack = append(stack, value)
				continue
			}
		}

		if !isDigit(value) && !isSpace(value) {
			if len(stack) > 0 {
				stack := revStack(stack)
				for i, v := range stack {
					topStPrior := prior(v)
					if topStPrior >= curPrior {
						outputStr = append(outputStr, v)
						outputStr = append(outputStr, ' ')
						stack = append(stack[:i], stack[i+1:]...)
					} else {
						stack = append(stack, v)
						stack = append(stack, ' ')
						stack = revStack(stack)
						break
					}
				}
				if len(stack) == 0 {
					stack = append(stack, value)
				}
			}
		}
	}

	outputStr = append(outputStr, reverseStack(stack)...)
	return string(outputStr)
}
