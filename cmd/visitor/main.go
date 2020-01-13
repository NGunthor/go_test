package main

import (
	"fmt"
	"strings"

	"github.com/NGunthor/go_test/pkg/patterns/visitor"
)

func main() {
	e := visitor.NewAddition(
		visitor.NewAddition(
			visitor.NewLiteral(2),
			visitor.NewLiteral(3), '*'),
		visitor.NewLiteral(1488), '+')
	sb := strings.Builder{}
	expVisitor := visitor.NewExpressionVisitor(&sb)
	e.Accept(expVisitor)
	fmt.Println(sb.String())
}
