package main

import (
	"fmt"
	"github.com/NGunthor/go_test/patterns/observer"
)

func main() {
	pub := observer.NewPublisher()
	a := observer.NewObserverA()
	pub.Attach(a)
	pub.Attach(observer.NewObserverB())
	pub.Notify()
	pub.Attach(observer.NewObserverB())
	pub.Notify()
	pub.Show()
	fmt.Println("============", a)
	pub.Unpin(a)
	pub.Notify()
	pub.Show()
}
