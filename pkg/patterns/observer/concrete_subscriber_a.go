package observer

type concreteObserverA struct {
	val int
}

// Update refreshes the state of a concrete observer (implements Observer interface)
func (s *concreteObserverA) Update() {
	s.val++
}

// NewObserverA ...
func NewObserverA(val int) Observer {
	return &concreteObserverA{val:val}
}
