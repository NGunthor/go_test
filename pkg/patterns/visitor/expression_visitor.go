package visitor

import (
	"fmt"
	"strings"
)

type expressionVisitor struct {
	sb *strings.Builder
}

func (e *expressionVisitor) VisitLiteral(literal Literal) {
	e.sb.WriteString(fmt.Sprintf("%g", literal.GetValue()))
}

func (e *expressionVisitor) VisitAddition(a Addition) {
	e.sb.WriteString("(")
	a.GetLeft().Accept(e)
	e.sb.WriteString("+")
	a.GetRight().Accept(e)
	e.sb.WriteString(")")
}

func NewExpressionVisitor(sb *strings.Builder) *expressionVisitor {
	return &expressionVisitor{sb: sb}
}
