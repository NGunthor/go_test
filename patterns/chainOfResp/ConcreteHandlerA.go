package chainOfResp

type ConcreteHandlerA struct {
	next Handler
}

func (c *ConcreteHandlerA) Request(message string) string {
	if message != "fail" && c.next != nil {
		return c.next.Request(message)
	}
	return "ConcreteHandlerA"
}

func (c *ConcreteHandlerA) SetNext(next Handler) {
	c.next = next
}

func NewConcreteHandlerA() Handler {
	return &ConcreteHandlerA{}
}
