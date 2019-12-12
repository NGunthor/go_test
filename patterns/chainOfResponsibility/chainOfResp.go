package chainOfResponsibility

type Handler interface {
	HandleRequest(message int) string
}

type a struct {

}

func (a *a)HandleRequest(message int) string {

}

func main() {
	a := a{}

}

type ConcreteHandler struct {
	current Handler
	next    *ConcreteHandler
}

func (c *ConcreteHandler) Add(handler *Handler) {
	cur := c
	if c.current == nil {
		c.current = *handler
	}
	c.current = *handler
}



func NewConcreteHandler() *ConcreteHandler {
	return &ConcreteHandler{current: nil, next: nil}
}
