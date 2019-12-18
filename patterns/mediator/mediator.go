package mediator

import (
	"log"
)

// Mediator provides
type Mediator interface {
	Notify(id rune, amount int)
}

type mediator struct {
	colleagueA Account
	colleagueB Account
}

// Notify processes an operation from the mediator (implement Mediator interface)
func (m *mediator) Notify(id rune, amount int) {
	if m.colleagueA == nil || m.colleagueB == nil {
		log.Println("mediator required")
		return
	}
	if id == 'A' {
		if m.colleagueB.GetSum()-amount < 0 {
			log.Println("insufficient funds B")
		} else {
			m.colleagueA.add(amount)
			m.colleagueB.withdraw(amount)
		}
	}
	if id == 'B' {
		if m.colleagueA.GetSum()-amount < 0 {
			log.Println("insufficient funds A")
			return
		} else {
			m.colleagueB.add(amount)
			m.colleagueA.withdraw(amount)
		}
	}
}

// NewMediator ...
func NewMediator(a Account, b Account) Mediator {
	return &mediator{
		colleagueA: a,
		colleagueB: b,
	}
}
