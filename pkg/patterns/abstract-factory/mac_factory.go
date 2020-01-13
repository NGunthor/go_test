package abstract_factory

import "github.com/NGunthor/go_test/pkg/patterns/abstract-factory/mac-elements"

type macFactory struct {
}

// CreateButton initializes mac button
func (m *macFactory) CreateButton() painter {
	return mac_elements.NewMacButton()
}

// CreateCheckbox initializes mac checkbox
func (m *macFactory) CreateCheckbox() painter {
	return mac_elements.NewMacCheckbox()
}

// NewMacFactory ...
func NewMacFactory() GUIFactory {
	return &macFactory{}
}
