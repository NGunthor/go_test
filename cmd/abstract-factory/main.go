package main

import (
	abstract_factory "github.com/NGunthor/go_test/pkg/patterns/abstract-factory"
)

func main() {
	factory := abstract_factory.NewWinFactory()
	app := abstract_factory.NewApplication(factory)
	app.CreateUI()
	app.PaintCheckbox()
}
