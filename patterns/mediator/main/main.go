package main

import "github.com/NGunthor/go_test/patterns/mediator"

func main() {
	a := mediator.NewAccount(1000)
	b := mediator.NewAccount(1000)
	med := mediator.NewMediator(a, b)
	a.Send('B', 500)
	a.SetMediator(med)
	b.SetMediator(med)
	a.Send('B', 500)
	a.Send('B', 700)
	a.Show()
	b.Show()
}
