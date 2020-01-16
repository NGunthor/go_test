package observer

import (
	"fmt"
)

type observer interface {
	Update()
}

type Publisher interface {
	Attach(observer observer)
	Unpin(observer observer)
	Notify()
	Show()
}

type publisher struct {
	observers []observer
}

// Attach pins an observer to the publisher (implements Publisher interface)
func (p *publisher) Attach(obs observer) {
	p.observers = append(p.observers, obs)
}

// Notify notifies all of the Publisher's observers (implements Publisher interface)
func (p *publisher) Notify() {
	for _, obs := range p.observers {
		obs.Update()
	}
}

// Show shows the state of the observers (implements Publisher interface)
func (p *publisher) Show() {
	for _, obs := range p.observers {
		fmt.Println(obs)
	}
}

// Unpin detaches an observer from the Publisher (implements Publisher interface)
func (p *publisher) Unpin(observer observer) {
	for i, obs := range p.observers {
		if obs == observer {
			p.observers = append(p.observers[:i], p.observers[i+1:]...)
		}
	}
}

// NewPublisher is a constructor for a Publisher
func NewPublisher() Publisher {
	return &publisher{}
}
