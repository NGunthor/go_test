package mac_factory

type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

type macFactory struct {
}

func (m *macFactory) CreateButton() Button {
	return &macButton{}
}

func (m *macFactory) CreateCheckbox() Checkbox {
	return &macCheckbox{}
}

func NewMacFactory() GUIFactory {
	return &macFactory{}
}
