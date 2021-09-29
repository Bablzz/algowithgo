package algo

import (
	"strconv"
	"strings"
)

// Your message is a string containing space separated words.
// You need to encrypt each word in the message using the following rules:
// The first letter must be converted to its ASCII code.
// The second letter must be switched with the last letter
// kata https://www.codewars.com/kata/5848565e273af816fb000449

func EncryptThis(text string) string {
	if len(text) == 0 {
		return ""
	}
	var str string

	strArr := strings.Fields(text)
	for _, v := range strArr {
		if len(v) == 1 {
			str += strconv.Itoa(int(v[0])) + " "
			continue
		}
		if len(v) == 2 {
			str += strconv.Itoa(int(v[0])) + string(v[1]) + " "
			continue
		}
		second := string(v[1])
		last := string(v[len(v)-1])
		str += strconv.Itoa(int(v[0])) + last + v[2:len(v)-1] + second + " "
	}
	return strings.TrimSuffix(str, " ")
}
