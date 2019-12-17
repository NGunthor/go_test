package observer

type concreteObserverB struct {
	val int
}

// Update refreshes the state of a concrete observer (implements Observer interface)
func (s *concreteObserverB) Update() {
	s.val += 10
}

// NewObserverB is a constructor for concreteObserverB
func NewObserverB() Observer {
	return &concreteObserverB{}
}
