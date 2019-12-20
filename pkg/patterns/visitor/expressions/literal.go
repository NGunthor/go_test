package expressions

type Literal interface {
	GetValue() float64
}

type literal struct {
	value float64
}

func (l *literal) Accept(visitor visitor) {
	visitor.VisitLiteral(l)
}

func (l *literal) GetValue() float64 {
	return l.value
}

func NewLiteral(val float64) *literal {
	return &literal{value: val}
}
