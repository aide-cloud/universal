package strings

import (
	"strings"
	"unicode/utf8"
)

// IsEmpty returns true if the string is empty.
func IsEmpty(s string) bool {
	return len(s) == 0
}

// IsEmptyTrimmed returns true if the string is empty or contains only whitespace.
func IsEmptyTrimmed(s string) bool {
	return LenTrimmed(s) == 0
}

// UTF8Len returns the number of runes in the string.
func UTF8Len(s string) int {
	return utf8.RuneCountInString(s)
}

// UTF8LenTrimmed returns the number of runes in the string after trimming whitespace.
func UTF8LenTrimmed(s string) int {
	return utf8.RuneCountInString(strings.TrimSpace(s))
}

// LenTrimmed returns the number of runes in the string after trimming whitespace.
func LenTrimmed(s string) int {
	return len(strings.TrimSpace(s))
}

// ToLower returns a copy of the string with all Unicode letters mapped to their lower case.
func ToLower(s string) string {
	return strings.ToLower(s)
}

// ToUpper returns a copy of the string with all Unicode letters mapped to their upper case.
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToTitle returns a copy of the string with all Unicode letters mapped to their title case.
func ToTitle(s string) string {
	return strings.ToTitle(s)
}
