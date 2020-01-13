package memento

// Memento provides memento interface
type Memento interface {
	GetState() (int, string)
}

type memento struct {
	val  int
	name string
}

// GetState extracts the memento's current state
func (m *memento) GetState() (int, string) {
	return m.val, m.name
}

// NewMemento ...
func NewMemento(val int, name string) Memento {
	return &memento{
		val:  val,
		name: name,
	}
}
