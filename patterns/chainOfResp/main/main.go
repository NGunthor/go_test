package main

import (
	"fmt"
	"github.com/NGunthor/go_test/patterns/chainOfResp"
)

func main() {
	a := chainOfResp.NewConcreteHandlerA()
	b := chainOfResp.NewConcreteHandlerB()
	a.SetNext(b)
	b.SetNext(chainOfResp.NewConcreteHandlerC())
	fmt.Println(a.Request(""))
}
