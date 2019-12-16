package observer

type concreteObserverA struct {
	val int
}

//Updates the state of a concrete observer ((implements Observer interface)
func (s *concreteObserverA) Update() {
	s.val++
}

//NewObserverA is a constructor for concreteObserverA
func NewObserverA() Observer {
	return &concreteObserverA{}
}
