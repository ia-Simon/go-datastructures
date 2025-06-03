package usecases

import (
	"errors"
	"regexp"
	"strings"

	"github.com/ia-Simon/go-datastrutures/structures/stack"
)

var operPrecedence = map[rune]int{
	'+': 0,
	'-': 0,
	'*': 1,
	'/': 1,
	'^': 2,
}

// InfixToPostfix converts math expressions in infix format to
// postfix format e.g. "a+b" -> "ab+".
//
// Currently supported symbols are:
//   - operators: "+-*/^"
//   - delimiters: "()"
func InfixToPostfix(infix string) (string, error) {
	var postfix strings.Builder
	s := stack.New[rune]()

	for _, rn := range infix {
		charType := characterType(rn)

		if charType == charTypeSpace {
			continue
		}

		if charType == charTypeOperand {
			postfix.WriteRune(rn)
			continue
		}

		if charType == charTypeOperator {
			for {
				stackRn, ok := s.Peek()
				if !ok || characterType(stackRn) == charTypeOpenParen || operPrecedence[rn] > operPrecedence[stackRn] {
					break
				}

				s.Pop()
				postfix.WriteRune(stackRn)
			}

			s.Push(rn)
			continue
		}

		if charType == charTypeOpenParen {
			s.Push(rn)
		}

		if charType == charTypeCloseParen {
			for {
				stackRn, ok := s.Pop()
				if !ok {
					return "", errors.New("unexpected \")\"; no previous \"(\" detected")
				}

				if characterType(stackRn) == charTypeOpenParen {
					break
				}
				postfix.WriteRune(stackRn)
			}
		}
	}

	for !s.IsEmpty() {
		stackRn, _ := s.Pop()

		if characterType(stackRn) == charTypeOpenParen {
			return "", errors.New("unmatched \"(\" detected")
		}
		postfix.WriteRune(stackRn)
	}

	return postfix.String(), nil
}

type charType string

const (
	charTypeOperand    charType = "Operand"
	charTypeOperator   charType = "Operator"
	charTypeOpenParen  charType = "OpenParen"
	charTypeCloseParen charType = "CloseParen"
	charTypeSpace      charType = "Space"
)

func characterType(rn rune) charType {
	switch true {
	case regexp.MustCompile(`\(`).MatchString(string(rn)):
		return charTypeOpenParen
	case regexp.MustCompile(`\)`).MatchString(string(rn)):
		return charTypeCloseParen
	case regexp.MustCompile(`\s`).MatchString(string(rn)):
		return charTypeSpace
	default:
		if _, ok := operPrecedence[rn]; ok {
			return charTypeOperator
		}
		return charTypeOperand
	}
}
