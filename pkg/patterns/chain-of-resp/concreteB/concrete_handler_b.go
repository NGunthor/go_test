package chain_of_resp

type handler interface {
	Request(string) string
}

type concreteHandlerB struct {
	next handler
}

//Handler interface implementation
func (c *concreteHandlerB) Request(message string) string {
	if message != "fail" && c.next != nil {
		return c.next.Request(message)
	}
	return "concreteHandlerB"
}

//NewConcreteHandlerB ...
func NewConcreteHandlerB(next handler) handler {
	return &concreteHandlerB{next: next}
}
