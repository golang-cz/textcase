package textcase

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

type parser int

const (
	_             parser = iota //   _$$_This is some text, OK?!
	idle                        // 1 ↑↑↑↑                  ↑   ↑
	firstAlphaNum               // 2     ↑    ↑  ↑    ↑     ↑
	alphaNum                    // 3      ↑↑↑  ↑  ↑↑↑  ↑↑↑   ↑
	delimiter                   // 4         ↑  ↑    ↑    ↑   ↑
)

func (s parser) next(r rune) parser {
	switch s {
	case idle:
		if isAlphaNum(r) {
			return firstAlphaNum
		}
	case firstAlphaNum:
		if isAlphaNum(r) {
			return alphaNum
		}
		return delimiter
	case alphaNum:
		if !isAlphaNum(r) {
			return delimiter
		}
	case delimiter:
		if isAlphaNum(r) {
			return firstAlphaNum
		}
		return idle
	}
	return s
}

func isAlphaNum(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r)
}

// Mark letter case changes, ie. "camelCaseTEXT" -> "camel_Case_TEXT".
func markLetterCaseChanges(input string) string {
	var b strings.Builder

	wasLetter := false
	countConsecutiveUpperLetters := 0

	for i := 0; i < len(input); {
		r, size := utf8.DecodeRuneInString(input[i:])
		i += size

		if unicode.IsLetter(r) {
			if wasLetter && countConsecutiveUpperLetters > 1 && !unicode.IsUpper(r) {
				b.WriteString("_")
			}
			if wasLetter && countConsecutiveUpperLetters == 0 && unicode.IsUpper(r) {
				b.WriteString("_")
			}
		}

		wasLetter = unicode.IsLetter(r)
		if unicode.IsUpper(r) {
			countConsecutiveUpperLetters++
		} else {
			countConsecutiveUpperLetters = 0
		}

		b.WriteRune(r)
	}

	return b.String()
}
