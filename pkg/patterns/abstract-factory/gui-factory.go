package abstract_factory

type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Button
}
