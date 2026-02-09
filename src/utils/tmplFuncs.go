package utils

import (
	"fmt"
	"strings"
	"time"
	"unicode"

	"github.com/fdanctl/p5r-stats/src/models"
)

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

// TimeToString return a date string in "Friday, 2 January, 2026" format.
func TimeToString(date time.Time) string {
	weekday := date.Weekday().String()
	month := date.Month().String()
	day := date.Day()
	year := date.Year()

	return fmt.Sprintf(
		"%s, %d %s, %d",
		weekday, day, month, year,
	)
}

func StatToString(s models.Stat) string {
	fmt.Println("stat", s.String())
	return s.String()
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
