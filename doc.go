// Package calculator provides a flexible and precise arithmetic expression evaluator.
//
// Overview:
// The calculator package allows parsing and evaluating mathematical expressions with support for
// variables, precision control, and detailed error handling. It is designed for applications
// requiring high decimal precision, such as financial or scientific calculations.
//
// Features:
// - Basic operations: +, -, *, /, and parentheses.
// - Variable substitution using key-value maps.
// - Precision control for both final results and intermediate calculations.
// - Comprehensive error reporting (syntax errors, division by zero, undefined variables).
//
// Usage:
// To evaluate an expression directly:
//
//	result, err := calculator.Evaluate("(x + y) * 2", map[string]decimal.Decimal{
//	    "x": decimal.NewFromInt(5),
//	    "y": decimal.NewFromInt(3),
//	})
//
// To compile an expression for reuse:
//
//	expr, err := calculator.Compile("a / b - c", calculator.WithFinalPrecision(2))
//	result, err := expr.Evaluate(vars)
//
// Precision Control Options:
// - Use WithFinalPrecision(n) to round the final result to n decimal places.
// - Use WithIntermediatePrecision(n) to prevent overflow during calculations.
//
// Error Handling:
// The package defines common error types for programmatic handling:
//
//	if errors.Is(err, calculator.ErrDivisionByZero) { ... }
//	if syntaxErr, ok := err.(*calculator.SyntaxError); ok { ... }
//
// Example:
// See examples in the package documentation or test files.
package calculator
