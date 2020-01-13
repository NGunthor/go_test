package visitor

// Expression is an interface for communicate expressions with visitor
type Expression interface {
	Accept(visitor Visitor)
}
