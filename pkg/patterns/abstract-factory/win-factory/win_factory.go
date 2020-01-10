package win_factory

type guiFactory interface {
	CreateButton() button
	CreateCheckbox() checkbox
}

type winFactory struct {
}

func (w *winFactory) CreateButton() button {
	return &winButton{}
}

func (w *winFactory) CreateCheckbox() checkbox {
	return &winCheckbox{}
}

func NewWinFactory() guiFactory {
	return &winFactory{}
}
