package main

import (
	"fmt"
	chainOfResp2 "github.com/NGunthor/go_test/pkg/patterns/chain-of-resp"
)

func main() {
	a := chainOfResp2.NewConcreteHandlerA(chainOfResp2.NewConcreteHandlerB(chainOfResp2.NewConcreteHandlerC(nil)))
	fmt.Println(a.Request("asd"))
}
