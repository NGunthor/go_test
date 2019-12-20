package visitor

import (
	"fmt"
	"strings"
)

type expression interface {
	Accept(visitor Visitor)
}

type literal interface {
	GetValue() expression
}

type addition interface {
	GetLeft() expression
	GetRight() expression
}

type expressionVisitor struct {
	sb *strings.Builder
}

func (e *expressionVisitor) VisitLiteral(literal literal) {
	e.sb.WriteString(fmt.Sprintf("%g", literal.GetValue()))
}

func (e *expressionVisitor) VisitAddition(a addition) {
	e.sb.WriteString("(")
	a.GetLeft().Accept(e)
	e.sb.WriteString("+")
	a.GetRight().Accept(e)
	e.sb.WriteString(")")
}

func NewExpressionVisitor(sb *strings.Builder) *expressionVisitor {
	return &expressionVisitor{sb: sb}
}
