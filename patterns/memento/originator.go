package memento

import "fmt"

type Originator interface {
	SetState(int, string)
	GetState() (int, string)
	SaveState() Memento
	BackUpState(Memento)
	Show()
}

type originator struct {
	val  int
	name string
}

func (o *originator) SaveState() Memento {
	return &memento{val: o.val, name: o.name}
}

func (o *originator) BackUpState(mem Memento) {
	o.val, o.name = mem.getState()
}

func (o *originator) SetState(val int, name string) {
	o.val, o.name = val, name
}

func (o *originator) GetState() (int, string) {
	return o.val, o.name
}

func (o *originator) Show() {
	fmt.Println(o.name, o.val)
}

func NewOriginator() Originator {
	return &originator{}
}
