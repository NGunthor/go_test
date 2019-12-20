package visitor

type Visitor interface {
	VisitLiteral(literal Literal)
	VisitAddition(addition Addition)
}
