package observer

import "fmt"

type concretePublisher struct {
	observers []Observer
}

// Attach pins an Observer to the publisher (implements Publisher interface)
func (p *concretePublisher) Attach(obs Observer) {
	p.observers = append(p.observers, obs)
}

// Notify notifies all of the Publisher's observers (implements Publisher interface)
func (p *concretePublisher) Notify() {
	for _, obs := range p.observers {
		obs.Update()
	}
}

// Show shows the state of the Observers (implements Publisher interface)
func (p *concretePublisher) Show() {
	for _, obs := range p.observers {
		fmt.Println(obs)
	}
}

// Unpin detaches an observer from the Publisher (implements Publisher interface)
func (p *concretePublisher) Unpin(observer Observer) {
	for i, obs := range p.observers {
		if obs == observer {
			p.observers = append(p.observers[:i], p.observers[i+1:]...)
		}
	}
}

// NewPublisher ...
func NewPublisher() Publisher {
	return &concretePublisher{}
}
