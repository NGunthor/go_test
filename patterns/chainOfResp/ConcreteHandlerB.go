package chainOfResp

type concreteHandlerB struct {
	next Handler
}

//Handler interface implementation
func (c *concreteHandlerB) Request(message string) string {
	if message != "fail" && c.next != nil {
		return c.next.Request(message)
	}
	return "concreteHandlerB"
}

//NewConcreteHandlerB ...
func NewConcreteHandlerB(next Handler) Handler {
	return &concreteHandlerB{next: next}
}
