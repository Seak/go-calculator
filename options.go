// calculator/options.go
package calculator

// Option 配置函数类型
type Option func(*options)

type options struct {
	finalPrecision        int32
	intermediatePrecision int32
}

var defaultOptions = options{-1, -1}

// WithFinalPrecision 设置最终结果的精度（小数位数）
// 示例：WithFinalPrecision(2) → 结果保留2位小数
func WithFinalPrecision(p int32) Option {
	return func(o *options) { o.finalPrecision = p }
}

// WithIntermediatePrecision 设置中间结果的精度（避免精度溢出）
// 示例：WithIntermediatePrecision(2) → 结果保留2位小数
func WithIntermediatePrecision(p int32) Option {
	return func(o *options) { o.intermediatePrecision = p }
}
