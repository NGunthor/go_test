package visitor

type Addition interface {
	GetLeft() Expression
	GetRight() Expression
}

type addition struct {
	left  Expression
	right Expression
}

func (a *addition) Accept(visitor Visitor) {
	visitor.VisitAddition(a)
}

func (a *addition) GetLeft() Expression {
	return a.left
}

func (a *addition) GetRight() Expression {
	return a.right
}

func NewAddition(left Expression, right Expression) *addition {
	return &addition{
		left:  left,
		right: right,
	}
}
