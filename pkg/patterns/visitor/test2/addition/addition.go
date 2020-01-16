package addition

import "github.com/NGunthor/go_test/pkg/patterns/visitor/test2/literal"

type visitor interface {
	VisitLiteral(literal literal.Literal)
	VisitAddition(addition Addition)
}

type Expression interface {
	Accept(visitor visitor)
}

// Addition provides interface for addition
type Addition interface {
	GetLeft() Expression
	GetRight() Expression
	GetOperation() rune
}

type addition struct {
	left      Expression
	right     Expression
	operation rune
}

// Accept extends addition's functional via Visitor (implements Expression interface)
func (a *addition) Accept(visitor visitor) {
	visitor.VisitAddition(a)
}

// GetLeft gets left part of Expression
func (a *addition) GetLeft() Expression {
	return a.left
}

// GetRight gets right part of Expression
func (a *addition) GetRight() Expression {
	return a.right
}

func (a *addition) GetOperation() rune {
	return a.operation
}

// NewAddition ...
func NewAddition(left Expression, right Expression, operation rune) Expression {
	return &addition{
		left:      left,
		right:     right,
		operation: operation,
	}
}
