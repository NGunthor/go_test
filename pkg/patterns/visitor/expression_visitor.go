package visitor

import (
	"fmt"
	"strings"
)

// Visitor provides interface for visitor functional
type Visitor interface {
	VisitLiteral(literal Literal)
	VisitAddition(addition Addition)
}

type expressionVisitor struct {
	sb *strings.Builder
}

// VisitLiteral extends literal's functional
func (e *expressionVisitor) VisitLiteral(literal Literal) {
	e.sb.WriteString(fmt.Sprintf("%g", literal.GetValue()))
}

// VisitAddition extends addition's functional
func (e *expressionVisitor) VisitAddition(a Addition) {
	e.sb.WriteString("(")
	a.GetLeft().Accept(e)
	e.sb.WriteString(string(a.GetOperation()))
	a.GetRight().Accept(e)
	e.sb.WriteString(")")
}

// NewExpressionVisitor
func NewExpressionVisitor(sb *strings.Builder) Visitor {
	return &expressionVisitor{sb: sb}
}
