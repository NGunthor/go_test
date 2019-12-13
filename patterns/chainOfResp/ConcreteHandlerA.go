package chainOfResp

type concreteHandlerA struct {
	next Handler
}

func (c *concreteHandlerA) Request(message string) string {
	if message != "fail" && c.next != nil {
		return c.next.Request(message)
	}
	return "ConcreteHandlerA"
}

func (c *concreteHandlerA) SetNext(next Handler) {
	c.next = next
}

func NewConcreteHandlerA() Handler {
	return &concreteHandlerA{}
}
