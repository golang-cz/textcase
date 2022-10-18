package textcase

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// Converts input string to "camelCase" (lower camel case) naming convention.
// Removes all whitespace and special characters. Supports Unicode characters.
func CamelCase(input string) string {
	var b strings.Builder

	state := idle
	for i := 0; i < len(input); {
		r, size := utf8.DecodeRuneInString(input[i:])
		i += size
		state = state.next(r)
		switch state {
		case firstAlphaNum:
			if b.Len() > 0 {
				b.WriteRune(unicode.ToUpper(r))
			} else {
				b.WriteRune(unicode.ToLower(r))
			}
		case alphaNum:
			b.WriteRune(unicode.ToLower(r))
		}
	}
	return b.String()
}
