package memento

// Memento provides memento interface
type Memento interface {
	getState() (int, string)
}

type memento struct {
	val  int
	name string
}

func (m *memento) getState() (int, string) {
	return m.val, m.name
}
