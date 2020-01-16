package concreteC

type handler interface {
	Request(string) string
}

type concreteHandlerC struct {
	next handler
}

//Handler interface implementation
func (c *concreteHandlerC) Request(message string) string {
	if message != "fail" && c.next != nil {
		return c.next.Request(message)
	}
	return "ConcreteHandlerC"
}

//NewConcreteHandlerC
func NewConcreteHandlerC(next handler) handler {
	return &concreteHandlerC{next: next}
}
