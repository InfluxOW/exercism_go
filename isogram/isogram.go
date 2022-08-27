package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	for i, l := range strings.ToLower(word) {
		if unicode.IsLetter(l) && strings.ContainsRune(word[i+1:], l) {
			return false
		}
	}
	return true
}
