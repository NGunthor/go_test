package chainOfResp

type concreteHandlerC struct {
	next Handler
}

func (c *concreteHandlerC) Request(message string) string {
	if message != "fail" && c.next != nil {
		return c.next.Request(message)
	}
	return "ConcreteHandlerC"
}

func (c *concreteHandlerC) SetNext(next Handler) {
	c.next = next
}

func NewConcreteHandlerC() Handler {
	return &concreteHandlerC{}
}
