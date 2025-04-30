// calculator/evaluate.go
package calculator

import (
	"github.com/shopspring/decimal"
)

// Evaluate 直接编译并求值表达式（快捷方法）
// 参数：
//   - expr: 数学表达式（支持加减乘除、括号、变量）
//   - vars: 变量映射表
//   - opts: 配置选项（如精度控制）
//
// 返回：
//   - decimal.Decimal: 计算结果
//   - error: 编译或运行错误
func Evaluate(expr string, vars map[string]decimal.Decimal, opts ...Option) (decimal.Decimal, error) {
	compiled, err := Compile(expr, opts...)
	if err != nil {
		return decimal.Zero, err
	}
	return compiled.Evaluate(vars)
}
