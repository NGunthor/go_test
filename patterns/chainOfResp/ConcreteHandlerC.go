package chainOfResp

type concreteHandlerC struct {
	next Handler
}

//Handler interface implementation
func (c *concreteHandlerC) Request(message string) string {
	if message != "fail" && c.next != nil {
		return c.next.Request(message)
	}
	return "ConcreteHandlerC"
}

//NewConcreteHandlerC
func NewConcreteHandlerC(next Handler) Handler {
	return &concreteHandlerC{next: next}
}
