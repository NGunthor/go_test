package main

import (
	"fmt"
	"github.com/NGunthor/go_test/pkg/patterns/proxy"
)

func main() {
	server := proxy.NewHttpServer(proxy.NewApplication())
	fmt.Println(server.HandleRequest("/myProfile", "GET"))
}
