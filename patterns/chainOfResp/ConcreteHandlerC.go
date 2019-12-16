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

func NewConcreteHandlerC(next *Handler) Handler {
	return &concreteHandlerC{next: *next}
}
