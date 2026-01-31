package utils

import (
	"fmt"
	"strings"
	"unicode"
)

// Upper returns a string with all letters converted to uppercase.
func Upper(s string) string {
	return strings.ToUpper(s)
}

// Capitalize returns a string with the first letter capitalized.
func Capitalize(s string) string {
	if s == "" {
		return s
	}

	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// ToTitleCase returns a string with all words capitalized.
func ToTitleCase(s string) string {
	words := strings.Fields(s)

	for i, w := range words {
		words[i] = Capitalize(w)
	}

	return strings.Join(words, " ")
}

// PadLeft formats a string with leading zeroes to make it at least `d` characters long.
func PadLeft(s string, d int) string {
	if len(s) >= d {
		return s
	}
	return PadLeft("0"+s, d)
}

// Dict creates a dictionary from an even number of key-value pairs.
func Dict(values ...any) (map[string]any, error) {
	if len(values)%2 != 0 {
		return nil, fmt.Errorf("dict expects even number of arguments")
	}
	m := make(map[string]any, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, fmt.Errorf("dict keys must be strings")
		}
		m[key] = values[i+1]
	}
	return m, nil
}
