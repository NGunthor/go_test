package chainOfResponsibility

import (
	"math/rand"
	"time"
)

type Handler interface {
	SendRequest() (string, bool)
}

type handler1 struct {

}

func (h *handler1)SendRequest() (string, bool) {
	return "I'm handler 1\n", NewBoolGen(rand.NewSource(time.Now().UnixNano())).Bool()
}

func NewHandler1() *handler1 {
	return &handler1{}
}

type handler2 struct {

}
func (h *handler2)SendRequest() (string, bool) {
	return "I'm handler 2\n", NewBoolGen(rand.NewSource(time.Now().UnixNano())).Bool()
}

func NewHandler2() *handler2 {
	return &handler2{}
}

type handler3 struct {

}

func (h *handler3)SendRequest() (string, bool) {
	return "I'm handler 3\n", NewBoolGen(rand.NewSource(time.Now().UnixNano())).Bool()
}

func NewHandler3() *handler3 {
	return &handler3{}
}