package abstract_factory

// GUIFactory provides interface for abstract factory
type GUIFactory interface {
	CreateButton() painter
	CreateCheckbox() painter
}
