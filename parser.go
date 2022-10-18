package textcase

import (
	"unicode"
)

type parserStateMachine int

const (
	_             parserStateMachine = iota //   _$$_This is some text, OK?!
	idle                                    // 1 ↑↑↑↑                  ↑   ↑
	firstAlphaNum                           // 2     ↑    ↑  ↑    ↑     ↑
	alphaNum                                // 3      ↑↑↑  ↑  ↑↑↑  ↑↑↑   ↑
	delimiter                               // 4         ↑  ↑    ↑    ↑   ↑
)

func (s parserStateMachine) next(r rune) parserStateMachine {
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
