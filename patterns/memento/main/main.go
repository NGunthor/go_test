package main

import (
	"fmt"

	"github.com/NGunthor/go_test/patterns/memento"
)

func main() {
	ct := memento.NewCareTaker()
	o := memento.NewOriginator()

	o.SetState(10, "slava")
	ct.PushMem(o.SaveState())
	o.Show()
	fmt.Println("--------")

	o.SetState(20, "asdasd")
	ct.PushMem(o.SaveState())
	o.Show()
	fmt.Println("--------")

	ct.Show()
	fmt.Println("--------")

	if mem, err := ct.GetMem(); err == nil {
		o.BackUpState(mem)
	}
	ct.Show()
	o.Show()
	if mem, err := ct.GetMem(); err == nil {
		o.BackUpState(mem)
	}
	o.Show()
	if mem, err := ct.GetMem(); err == nil {
		o.BackUpState(mem)
	}
	o.Show()
	ct.Show()
}
