package main

import (
	mediator2 "github.com/NGunthor/go_test/pkg/patterns/mediator"
)

func main() {
	a := mediator2.NewAccount(1000)
	b := mediator2.NewAccount(1000)
	med := mediator2.NewMediator(a, b)
	a.Send('B', 500)
	a.SetMediator(med)
	b.SetMediator(med)
	a.Send('B', 500)
	a.Send('B', 700)
	a.Show()
	b.Show()
}
