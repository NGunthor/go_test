package main

import (
	handlerA "github.com/NGunthor/go_test/pkg/patterns/chain-of-resp/concreteA"
	handlerB "github.com/NGunthor/go_test/pkg/patterns/chain-of-resp/concreteB"
	handlerC "github.com/NGunthor/go_test/pkg/patterns/chain-of-resp/concreteC"
)

func main() {
	a := handlerA.NewConcreteHandlerA(handlerB.NewConcreteHandlerB(handlerC.NewConcreteHandlerC(nil)))
	a.Request("fail")
}
