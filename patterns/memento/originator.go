package memento

import "fmt"

// Originator provides originator interface
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

// SaveState saves state of originator into Memento interface (implements Originator interface)
func (o *originator) SaveState() Memento {
	return &memento{val: o.val, name: o.name}
}

// BackUpState gets an Memento and backs up by updating itself (implements Originator interface)
func (o *originator) BackUpState(mem Memento) {
	o.val, o.name = mem.getState()
}

// SetState sets its state by input
func (o *originator) SetState(val int, name string) {
	o.val, o.name = val, name
}

// GetState gets originator's state (implements Originator interface)
func (o *originator) GetState() (int, string) {
	return o.val, o.name
}
// Shows shows state of originator struct
func (o *originator) Show() {
	fmt.Println(o.name, o.val)
}

// NewOriginator ...
func NewOriginator() Originator {
	return &originator{}
}
