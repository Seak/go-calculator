// calculator/errors.go
package calculator

import "fmt"

// 定义错误类型
var (
	ErrDivisionByZero    = fmt.Errorf("division by zero")
	ErrUndefinedVariable = fmt.Errorf("undefined variable")
	ErrInvalidExpression = fmt.Errorf("invalid expression")
	ErrUnknownFunction = fmt.Errorf("unknown function")
)

// SyntaxError 表示表达式语法错误
type SyntaxError struct {
	Position int
	Message  string
}

func (e *SyntaxError) Error() string {
	return fmt.Sprintf("syntax error at position %d: %s", e.Position, e.Message)
}
