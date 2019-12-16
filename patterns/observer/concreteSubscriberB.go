package observer

type concreteObserverB struct {
	val int
}

//Updates the state of a concrete observer
func (s *concreteObserverB) Update() {
	s.val += 10
}

//NewObserverB is a constructor for concreteObserverB
func NewObserverB() Observer {
	return &concreteObserverB{}
}
