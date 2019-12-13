package chainOfResp

type concreteHandlerB struct {
	next Handler
}

func (c *concreteHandlerB)Request(message string) string{
	if message != "fail" && c.next != nil{
		return c.next.Request(message)
	}
	return "concreteHandlerB"
}

func (c *concreteHandlerB)SetNext(next Handler) {
	c.next = next
}

func NewConcreteHandlerB() Handler {
	return &concreteHandlerB{}
}