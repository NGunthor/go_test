package visitor

import (
	"fmt"
	"strings"

	"github.com/NGunthor/go_test/pkg/patterns/visitor/test2/addition"
	"github.com/NGunthor/go_test/pkg/patterns/visitor/test2/literal"
)

// Visitor provides interface for visitor functional
type Visitor interface {
	VisitLiteral(literal literal.Literal)
	VisitOperation(addition addition.Addition)
}

type expressionVisitor struct {
	sb *strings.Builder
}

// VisitLiteral extends literal's functional
func (e *expressionVisitor) VisitLiteral(literal literal.Literal) {
	e.sb.WriteString(fmt.Sprintf("%g", literal.GetValue()))
}

// VisitAddition extends addition's functional
func (e *expressionVisitor) VisitOperation(a addition.Addition) {
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
