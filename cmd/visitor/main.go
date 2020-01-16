package main

import (
	"fmt"
	"strings"

	"github.com/NGunthor/go_test/pkg/patterns/visitor/test/algebra"
	"github.com/NGunthor/go_test/pkg/patterns/visitor/test/visitor"
)

func main() {
	e := algebra.NewOperation(
		algebra.NewOperation(
			algebra.NewLiteral(2),
			algebra.NewLiteral(3), '-'),
		algebra.NewLiteral(1488), '+')
	sb := strings.Builder{}
	expVisitor := visitor.NewExpressionVisitor(&sb)
	e.Accept(expVisitor)
	fmt.Println(sb.String())
}
