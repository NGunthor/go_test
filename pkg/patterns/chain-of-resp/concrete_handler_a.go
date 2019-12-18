package chain_of_resp

type concreteHandlerA struct {
	next Handler
}

//Handler interface implementation
func (c *concreteHandlerA) Request(message string) string {
	if message != "fail" && c.next != nil {
		return c.next.Request(message)
	}
	return "ConcreteHandlerA"
}

//NewConcreteHandlerA ...
func NewConcreteHandlerA(next Handler) Handler {
	return &concreteHandlerA{next: next}
}
