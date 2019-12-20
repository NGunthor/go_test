package visitor

type Visitor interface {
	VisitLiteral(literal literal)
	VisitAddition(addition addition)
}
