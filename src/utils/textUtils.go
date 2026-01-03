package utils

import (
	"strings"
	"unicode"
)

func Upper(s string) string {
	return strings.ToUpper(s)
}

func Capitalize(s string) string {
	if s == "" {
		return s
	}

	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func ToTitleCase(s string) string {
	words := strings.Fields(s)

	for i, w := range words {
		words[i] = Capitalize(w)
	}

	return strings.Join(words, " ")
}
