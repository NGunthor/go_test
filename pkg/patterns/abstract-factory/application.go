package abstract_factory

type Application interface {
	CreateUI()
	PaintButton()
	PaintCheckbox()

}

type application struct {
	factory  GUIFactory
	button   Button
	checkbox Checkbox
}

func (a *application) CreateUI() {
	a.button = a.factory.CreateButton()
	a.checkbox = a.factory.CreateCheckbox()
}

func (a *application) PaintButton() {
	a.button.Paint()
}

func (a *application) PaintCheckbox() {
	a.checkbox.Paint()
}

func NewApplication(factory GUIFactory) Application {
	return &application{factory: factory}
}
