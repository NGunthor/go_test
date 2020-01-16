package main

import (
	"math/rand"
	"time"

	"github.com/NGunthor/go_test/pkg/patterns/proxy"
	bool_generator "github.com/NGunthor/go_test/pkg/patterns/proxy/bool-generator"
	httpserver "github.com/NGunthor/go_test/pkg/patterns/proxy/http-server"
)

func main() {
	generator := bool_generator.NewBoolGenerator(rand.NewSource(time.Now().UnixNano()))
	httpServer := httpserver.NewHttpServer()
	validator := proxy.NewValidator(httpServer, generator)
	validator.HandleRequest("", "")
}
