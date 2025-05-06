// calculator/token.go
package calculator

import "github.com/shopspring/decimal"

type tokenType int

const (
	tokenNumber tokenType = iota
	tokenAdd
	tokenSub
	tokenMul
	tokenDiv
	tokenLeftParen
	tokenRightParen
	tokenIdent
	tokenComma
	tokenEOF
)

type token struct {
	typ   tokenType
	value decimal.Decimal
	ident string
	pos   int
}
