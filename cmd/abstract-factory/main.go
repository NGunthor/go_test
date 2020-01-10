package main

import (
	abstract_factory "github.com/NGunthor/go_test/pkg/patterns/abstract-factory"
	mac_factory "github.com/NGunthor/go_test/pkg/patterns/abstract-factory/mac-factory"
	win_factory "github.com/NGunthor/go_test/pkg/patterns/abstract-factory/win-factory"
)

func main() {
	factory := mac_factory.NewMacFactory()
	//app := abstract_factory.NewApplication(factory)
}
