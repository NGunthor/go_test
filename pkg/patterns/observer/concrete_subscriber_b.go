package observer

type concreteObserverB struct {
	val int
}

// Update refreshes the state of a concrete observer (implements Observer interface)
func (s *concreteObserverB) Update() {
	s.val += 10
}

// NewObserverB ...
func NewObserverB(val int) Observer {
	return &concreteObserverB{val:val}
}
