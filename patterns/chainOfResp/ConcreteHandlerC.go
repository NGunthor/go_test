package chainOfResp

type ConcreteHandlerC struct {
	next Handler
}

func (c *ConcreteHandlerC)Request(message string) string{
	if message != "fail" && c.next != nil{
		return c.next.Request(message)
	}
	return "ConcreteHandlerC"
}

func (c *ConcreteHandlerC)SetNext(next Handler) {
	c.next = next
}

func NewConcreteHandlerC() Handler {
	return &ConcreteHandlerC{}
}
