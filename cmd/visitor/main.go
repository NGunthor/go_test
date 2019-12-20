package main

import (
	"fmt"
	"github.com/NGunthor/go_test/pkg/patterns/visitor"
	"github.com/NGunthor/go_test/pkg/patterns/visitor/expressions"
	"strings"
)

func main() {
	e := expressions.NewAddition(
		expressions.NewAddition(
			expressions.NewLiteral(1),
			expressions.NewLiteral(2)),
		expressions.NewLiteral(3))
	sb := strings.Builder{}
	expVisitor := visitor.NewExpressionVisitor(&sb)
	e.Accept(expVisitor)
	fmt.Println(sb.String())
}
