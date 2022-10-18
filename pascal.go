package textcase

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// Converts input string to "PascalCase" (upper camel case) naming convention.
// Removes all whitespace and special characters. Supports Unicode characters.
func PascalCase(input string) string {
	var b strings.Builder

	stateMachine := idle
	for i := 0; i < len(input); {
		r, size := utf8.DecodeRuneInString(input[i:])
		i += size
		stateMachine = stateMachine.next(r)
		switch stateMachine {
		case firstAlphaNum:
			b.WriteRune(unicode.ToUpper(r))
		case alphaNum:
			b.WriteRune(unicode.ToLower(r))
		}
	}
	return b.String()
}
