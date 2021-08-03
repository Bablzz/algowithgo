package algo

import "unicode"

func IsPalindrome(in string) bool {
	runes := make([]rune, 0, len(in))

	for _, r := range in {
		if !unicode.IsLetter(r) {
			continue
		}
		runes = append(runes, unicode.ToLower(r))
	}

	for i := range runes {
		if runes[i] != runes[len(runes)-i-1] {
			return false
		}
	}
	return true
}
