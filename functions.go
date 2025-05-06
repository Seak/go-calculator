package calculator

import (
	"fmt"
	"sync"

	"github.com/shopspring/decimal"
)

var (
	funcsMu sync.RWMutex
	funcs   = make(map[string]func(...decimal.Decimal) (decimal.Decimal, error))
)

// RegisterFunction 注册自定义函数
// 参数：
//   - name: 函数名（需唯一）
//   - fn: 函数实现，接收可变参数并返回结果或错误
func RegisterFunction(name string, fn func(...decimal.Decimal) (decimal.Decimal, error)) error {
	funcsMu.Lock()
	defer funcsMu.Unlock()
	if _, exists := funcs[name]; exists {
		return fmt.Errorf("function '%s' already registered", name)
	}
	funcs[name] = fn
	return nil
}

func getFunction(name string) (func(...decimal.Decimal) (decimal.Decimal, error), bool) {
	funcsMu.RLock()
	defer funcsMu.RUnlock()
	fn, exists := funcs[name]
	return fn, exists
}
