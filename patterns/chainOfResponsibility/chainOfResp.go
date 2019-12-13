package chainOfResponsibility

import "strconv"

type ConcreteHandler struct {
	current Handler
	next    *ConcreteHandler
}

func (c *ConcreteHandler) Add(handler Handler) {
	cur := c

	if cur.current == nil {
		cur.current = handler
		return
	}
	for cur.next != nil {
		cur = cur.next
	}
	newHandler := NewConcreteHandler()
	newHandler.current = handler
	cur.next = newHandler

}

func (c *ConcreteHandler) HandleRequest() (result string) {
	cur := c
	for i := 1; cur != nil; i++ {
		if response, done := cur.current.SendRequest(); done {
			result += response
		} else {
			return result + "Bad response on " + strconv.Itoa(i) + "\n"
		}
		cur = cur.next
	}
	return
}

func NewConcreteHandler() *ConcreteHandler {
	return &ConcreteHandler{current: nil, next: nil}
}
