package abstract_factory

import win_factory "github.com/NGunthor/go_test/pkg/patterns/abstract-factory/win-elements"

type winFactory struct {
}

// CreateButton initializes windows button
func (w *winFactory) CreateButton() painter {
	return win_factory.NewWinButton()
}

// CreateCheckbox initializes windows checkbox
func (w *winFactory) CreateCheckbox() painter {
	return win_factory.NewWinCheckbox()
}

// NewWinFactory ...
func NewWinFactory() GUIFactory {
	return &winFactory{}
}
