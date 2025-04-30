// calculator/ast.go
package calculator

import (
	"fmt"

	"github.com/shopspring/decimal"
)

// 定义AST节点接口和具体实现
type node interface {
	eval(vars map[string]decimal.Decimal, opts options) (decimal.Decimal, error)
}

type binaryNode struct {
	op    tokenType
	left  node
	right node
}

type numberNode struct {
	value decimal.Decimal
}

type variableNode struct {
	name string
}

type unaryNode struct {
	op   tokenType
	expr node
}

func (n *numberNode) eval(_ map[string]decimal.Decimal, _ options) (decimal.Decimal, error) {
	return n.value, nil
}

func (n *variableNode) eval(vars map[string]decimal.Decimal, _ options) (decimal.Decimal, error) {
	val, exists := vars[n.name]
	if !exists {
		return decimal.Zero, fmt.Errorf("undefined variable '%s'", n.name)
	}
	return val, nil
}

func (n *unaryNode) eval(vars map[string]decimal.Decimal, opts options) (decimal.Decimal, error) {
	val, err := n.expr.eval(vars, opts)
	if err != nil {
		return decimal.Zero, err
	}
	if n.op == tokenSub {
		return val.Neg(), nil
	}
	return val, nil
}

func (n *binaryNode) eval(vars map[string]decimal.Decimal, opts options) (decimal.Decimal, error) {
	leftVal, err := n.left.eval(vars, opts)
	if err != nil {
		return decimal.Zero, err
	}

	rightVal, err := n.right.eval(vars, opts)
	if err != nil {
		return decimal.Zero, err
	}

	var result decimal.Decimal
	switch n.op {
	case tokenAdd:
		result = leftVal.Add(rightVal)
	case tokenSub:
		result = leftVal.Sub(rightVal)
	case tokenMul:
		result = leftVal.Mul(rightVal)
	case tokenDiv:
		if rightVal.IsZero() {
			return decimal.Zero, fmt.Errorf("division by zero")
		}
		result = leftVal.Div(rightVal)
	default:
		return decimal.Zero, fmt.Errorf("unknown operator")
	}

	if opts.intermediatePrecision >= 0 {
		result = result.Round(opts.intermediatePrecision)
	}
	return result, nil
}
