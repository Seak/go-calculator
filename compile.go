// calculator/compile.go
package calculator

import (
	"fmt"

	"github.com/shopspring/decimal"
)

// CompiledExpression 表示编译后的表达式，用户通过此对象进行求值
type CompiledExpression struct {
	root node
	opts options
}

// Compile 将字符串表达式编译为可执行的表达式对象
// 参数：
//   - expr: 数学表达式（支持加减乘除、括号、变量）
//   - opts: 配置选项（如精度控制）
//
// 返回：
//   - *CompiledExpression: 编译后的表达式
//   - error: 语法错误或配置错误
func Compile(expr string, opts ...Option) (*CompiledExpression, error) {
	cfg := defaultOptions
	for _, opt := range opts {
		opt(&cfg)
	}

	p, err := newParser(expr, cfg)
	if err != nil {
		return nil, err
	}

	root, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	if p.currentToken.typ != tokenEOF {
		return nil, fmt.Errorf("unexpected token at position %d", p.currentToken.pos)
	}

	return &CompiledExpression{
		root: root,
		opts: cfg,
	}, nil
}

// Evaluate 执行编译后的表达式并返回结果
// 参数：
//   - vars: 变量映射表（变量名 → 值）
//
// 返回：
//   - decimal.Decimal: 计算结果
//   - error: 运行错误（如除零、未定义变量）
func (c *CompiledExpression) Evaluate(vars map[string]decimal.Decimal) (decimal.Decimal, error) {
	result, err := c.root.eval(vars, c.opts)
	if err != nil {
		return decimal.Zero, err
	}

	if c.opts.finalPrecision >= 0 {
		result = result.Round(c.opts.finalPrecision)
	}
	return result, nil
}
