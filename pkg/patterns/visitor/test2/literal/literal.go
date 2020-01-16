package literal

import "github.com/NGunthor/go_test/pkg/patterns/visitor/test2/addition"

type Expression interface {
	Accept(visitor visitor.Visitor)
}

// Literal provides interface for literal
type Literal interface {
	GetValue() float64
}

type literal struct {
	value float64
}

// Accept extends literal's functional via Visitor (implements Expression interface)
func (l *literal) Accept(visitor visitor.Visitor) {
	visitor.VisitLiteral(l)
}

// GetValue gets literal's value
func (l *literal) GetValue() float64 {
	return l.value
}

// NewLiteral ...
func NewLiteral(val float64) Expression {
	return &literal{value: val}
}
