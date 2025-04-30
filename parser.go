// calculator/parser.go
package calculator

import (
	"fmt"
)

type parser struct {
	lexer        *lexer
	currentToken token
	opts         options
}

func newParser(input string, opts options) (*parser, error) {
	l := newLexer(input)
	firstToken, err := l.nextToken()
	if err != nil {
		return nil, err
	}
	return &parser{
		lexer:        l,
		currentToken: firstToken,
		opts:         opts,
	}, nil
}

func (p *parser) eat(typ tokenType) error {
	if p.currentToken.typ == typ {
		tok, err := p.lexer.nextToken()
		if err != nil {
			return err
		}
		p.currentToken = tok
		return nil
	}
	return fmt.Errorf("syntax error at position %d: expected %v, got %v",
		p.currentToken.pos, typ, p.currentToken.typ)
}

func (p *parser) parseExpression() (node, error) {
	node, err := p.parseTerm()
	if err != nil {
		return nil, err
	}

	for {
		switch p.currentToken.typ {
		case tokenAdd, tokenSub:
			op := p.currentToken.typ
			if err := p.eat(op); err != nil {
				return nil, err
			}

			right, err := p.parseTerm()
			if err != nil {
				return nil, err
			}

			node = &binaryNode{
				op:    op,
				left:  node,
				right: right,
			}
		default:
			return node, nil
		}
	}
}

func (p *parser) parseTerm() (node, error) {
	node, err := p.parseFactor()
	if err != nil {
		return nil, err
	}

	for {
		switch p.currentToken.typ {
		case tokenMul, tokenDiv:
			op := p.currentToken.typ
			if err := p.eat(op); err != nil {
				return nil, err
			}

			right, err := p.parseFactor()
			if err != nil {
				return nil, err
			}

			node = &binaryNode{
				op:    op,
				left:  node,
				right: right,
			}
		default:
			return node, nil
		}
	}
}

func (p *parser) parseFactor() (node, error) {
	if p.currentToken.typ == tokenAdd || p.currentToken.typ == tokenSub {
		op := p.currentToken.typ
		if err := p.eat(op); err != nil {
			return nil, err
		}

		expr, err := p.parseFactor()
		if err != nil {
			return nil, err
		}

		return &unaryNode{op: op, expr: expr}, nil
	}

	switch p.currentToken.typ {
	case tokenNumber:
		val := p.currentToken.value
		if err := p.eat(tokenNumber); err != nil {
			return nil, err
		}
		return &numberNode{value: val}, nil
	case tokenIdent:
		name := p.currentToken.ident
		if err := p.eat(tokenIdent); err != nil {
			return nil, err
		}
		return &variableNode{name: name}, nil
	case tokenLeftParen:
		if err := p.eat(tokenLeftParen); err != nil {
			return nil, err
		}
		expr, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		if err := p.eat(tokenRightParen); err != nil {
			return nil, fmt.Errorf("missing closing parenthesis at position %d", p.currentToken.pos)
		}
		return expr, nil
	default:
		return nil, fmt.Errorf("unexpected token at position %d", p.currentToken.pos)
	}
}
