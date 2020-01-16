package main

import (
	"fmt"
	publisher "github.com/NGunthor/go_test/pkg/patterns/observer"
	"github.com/NGunthor/go_test/pkg/patterns/observer/subscriber1"
	"github.com/NGunthor/go_test/pkg/patterns/observer/subscriber2"
)

func main() {
	pub := publisher.NewPublisher()
	a := subscriber1.NewObserverA()
	pub.Attach(a)
	pub.Attach(subscriber2.NewObserverB())
	pub.Notify()
	pub.Attach(subscriber2.NewObserverB())
	pub.Notify()
	pub.Show()
	fmt.Println("============\n", a)
	pub.Unpin(a)
	pub.Notify()
	pub.Show()
}
