package main

import (
	"fmt"
	"github.com/NGunthor/go_test/patterns/chainOfResp"
)

func main() {
	a := chainOfResp.NewConcreteHandlerA(chainOfResp.NewConcreteHandlerB(chainOfResp.NewConcreteHandlerC(nil)))
	fmt.Println(a.Request("asd"))
}
