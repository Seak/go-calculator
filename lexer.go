// calculator/lexer.go
package calculator

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/shopspring/decimal"
)

var identRegex = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*`)

type lexer struct {
	input string
	pos   int
}

func newLexer(input string) *lexer {
	return &lexer{
		input: strings.TrimSpace(input),
		pos:   0,
	}
}

func (l *lexer) nextToken() (token, error) {
	for l.pos < len(l.input) && unicode.IsSpace(rune(l.input[l.pos])) {
		l.pos++
	}

	startPos := l.pos
	if l.pos >= len(l.input) {
		return token{typ: tokenEOF, pos: startPos}, nil
	}

	if matches := identRegex.FindString(l.input[l.pos:]); matches != "" {
		t := token{
			typ:   tokenIdent,
			ident: matches,
			pos:   startPos,
		}
		l.pos += len(matches)
		return t, nil
	}

	ch := l.input[l.pos]

	if unicode.IsDigit(rune(ch)) || ch == '.' {
		hasDot := ch == '.'
		l.pos++

		for l.pos < len(l.input) {
			c := l.input[l.pos]
			if unicode.IsDigit(rune(c)) {
				l.pos++
			} else if c == '.' {
				if hasDot {
					return token{}, fmt.Errorf("invalid number at position %d", startPos)
				}
				hasDot = true
				l.pos++
			} else {
				break
			}
		}

		numStr := l.input[startPos:l.pos]
		value, err := decimal.NewFromString(numStr)
		if err != nil {
			return token{}, fmt.Errorf("invalid number '%s' at position %d", numStr, startPos)
		}
		return token{typ: tokenNumber, value: value, pos: startPos}, nil
	}

	switch ch {
	case '+':
		l.pos++
		return token{typ: tokenAdd, pos: startPos}, nil
	case '-':
		l.pos++
		return token{typ: tokenSub, pos: startPos}, nil
	case '*':
		l.pos++
		return token{typ: tokenMul, pos: startPos}, nil
	case '/':
		l.pos++
		return token{typ: tokenDiv, pos: startPos}, nil
	case '(':
		l.pos++
		return token{typ: tokenLeftParen, pos: startPos}, nil
	case ')':
		l.pos++
		return token{typ: tokenRightParen, pos: startPos}, nil
	case ',':
        l.pos++
        return token{typ: tokenComma, pos: startPos}, nil
	default:
		return token{}, fmt.Errorf("invalid character '%c' at position %d", ch, startPos)
	}
}
