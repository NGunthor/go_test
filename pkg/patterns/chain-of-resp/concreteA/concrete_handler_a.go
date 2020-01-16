package chain_of_resp

type handler interface {
	Request(string) string
}

type concreteHandlerA struct {
	next handler
}

//Handler interface implementation
func (c *concreteHandlerA) Request(message string) string {
	if message != "fail" && c.next != nil {
		return c.next.Request(message)
	}
	return "ConcreteHandlerA"
}

//NewConcreteHandlerA ...
func NewConcreteHandlerA(next handler) handler {
	return &concreteHandlerA{next: next}
}
