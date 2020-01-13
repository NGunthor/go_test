package main

import (
	"fmt"

	memento2 "github.com/NGunthor/go_test/pkg/patterns/memento"
)

func main() {
	ct := memento2.NewCareTaker()
	o := memento2.NewOriginator()

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
