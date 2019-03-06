package camelsnake

import (
	"unicode"
)

type stateMachine int

const (
	idle          stateMachine = iota // 0 _$$_This is some text, OK?!
	firstAlphaNum                     // 1     ↑    ↑  ↑    ↑     ↑
	alphaNum                          // 2      ↑↑↑  ↑  ↑↑↑  ↑↑↑   ↑
	delimiter                         // 3         ↑  ↑    ↑    ↑   ↑
)

func (s stateMachine) next(r rune) stateMachine {
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
