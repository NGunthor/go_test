package visitor

type Expression interface {
	Accept(visitor Visitor)
}

