package main

import (
	"github.com/NGunthor/go_test/pkg/patterns/adapter"
	"github.com/NGunthor/go_test/pkg/patterns/adapter/transports"
)

func main() {
	driver := adapter.NewDriver()
	auto := transports.NewAuto()
	driver.Travel(auto)
	horse := transports.NewHorse()
	horseTransport := adapter.NewAnimalToTransportAdapter(horse)
	driver.Travel(horseTransport)
}
