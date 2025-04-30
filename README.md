# Precision Calculator in Go

高精度数学表达式计算器，支持变量替换和灵活精度控制，适用于财务、科学等需要精确计算的场景。

## 特性

- **基础运算**: `+` `-` `*` `/` 及括号优先级
- **变量支持**: 通过键值对注入变量值
- **精度控制**: 可配置最终结果和中间计算的精度
- **严格错误处理**: 语法错误、除零、未定义变量等明确错误类型
- **高性能**: 编译表达式复用，避免重复解析

## 安装

```bash
go get github.com/shopspring/decimal
go get github.com/yourusername/calculator
```

## 快速开始
```bash
package main

import (
  "fmt"
  "github.com/shopspring/decimal"
  "github.com/yourusername/calculator"
)

func main() {
  // 直接求值
  result, _ := calculator.Evaluate("(5.2 + 3)*x", map[string]decimal.Decimal{
    "x": decimal.NewFromInt(4),
  })
  fmt.Println(result) // 33.6

  // 编译后重复使用
  expr, _ := calculator.Compile("a/(b+c)", calculator.WithFinalPrecision(2))
  
  vars := map[string]decimal.Decimal{
    "a": decimal.NewFromFloat(10.5),
    "b": decimal.NewFromInt(3),
    "c": decimal.NewFromInt(2),
  }
  res, _ := expr.Evaluate(vars)
  fmt.Println(res) // 2.10
}
```

## 配置选项
|方法|说明|示例|
|----|----|----|
|WithFinalPrecision(n)|结果保留n位小数|Compile("x/3", WithFinalPrecision(4)) → 3.3333|
|WithIntermediatePrecision(n)|中间计算时限制精度|防止超大数溢出|
```bash
// 同时配置两种精度
expr, _ := calculator.Compile("(a + b)^10",
  calculator.WithFinalPrecision(2),
  calculator.WithIntermediatePrecision(8),
)
```

## 错误处理
```bash
result, err := calculator.Evaluate("5 / (y - y)", map[string]decimal.Decimal{"y": decimal.NewFromInt(3)})
if errors.Is(err, calculator.ErrDivisionByZero) {
  fmt.Println("捕获除零错误！")
}

if syntaxErr, ok := err.(*calculator.SyntaxError); ok {
  fmt.Printf("语法错误位置 %d: %s", syntaxErr.Position, syntaxErr.Message)
}
```

# 示例场景
## 财务计算
```bash
// 含税金额计算
expr, _ := calculator.Compile("price * (1 + tax_rate)")
result, _ := expr.Evaluate(map[string]decimal.Decimal{
  "price":    decimal.NewFromInt(2999),
  "tax_rate": decimal.NewFromFloat(0.08),
})
```

## 科学公式
```bash
// 动能公式 (1/2 mv²)
expr, _ := calculator.Compile("0.5 * m * v^2")
res, _ := expr.Evaluate(map[string]decimal.Decimal{
  "m": decimal.NewFromFloat(2.5),
  "v": decimal.NewFromFloat(10),
}) // 125
```
