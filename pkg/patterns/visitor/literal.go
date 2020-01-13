package visitor

// Literal provides interface for literal
type Literal interface {
	GetValue() float64
}

type literal struct {
	value float64
}

// Accept extends literal's functional via Visitor (implements Expression interface)
func (l *literal) Accept(visitor Visitor) {
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
