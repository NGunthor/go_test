package main

import (
	"fmt"
	observer2 "github.com/NGunthor/go_test/pkg/patterns/observer"
)

func main() {
	pub := observer2.NewPublisher()
	a := observer2.NewObserverA(10)
	pub.Attach(a)
	pub.Attach(observer2.NewObserverB(20))
	pub.Notify()
	pub.Attach(observer2.NewObserverB(30))
	pub.Notify()
	pub.Show()
	fmt.Println("============", a)
	pub.Unpin(a)
	pub.Notify()
	pub.Show()
}
