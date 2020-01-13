package abstract_factory

type painter interface {
	Paint()
}

// Application provides application's interface
type Application interface {
	CreateUI()
	PaintButton()
	PaintCheckbox()
}

type application struct {
	factory  GUIFactory
	button   painter
	checkbox painter
}

// CreateUI initializes application's UI elements
func (a *application) CreateUI() {
	a.button = a.factory.CreateButton()
	a.checkbox = a.factory.CreateCheckbox()
}

// PaintButton paints concrete button
func (a *application) PaintButton() {
	a.button.Paint()
}

// PaintCheckbox paints concrete checkbox
func (a *application) PaintCheckbox() {
	a.checkbox.Paint()
}

// NewApplication ...
func NewApplication(factory GUIFactory) Application {
	return &application{factory: factory}
}
