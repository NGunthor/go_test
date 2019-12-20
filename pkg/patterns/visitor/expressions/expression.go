package expressions

type Expression interface {
	Accept(visitor visitor)
}

type visitor interface {
	VisitLiteral(literal Literal)
	VisitAddition(addition Addition)
}
