package mediator

import (
	"fmt"
	"log"
)

// Account provides an account interface
type Account interface {
	add(int)
	withdraw(int)
	Send(rune, int)
	GetSum() int
	SetMediator(Mediator)
	Show()
}

type account struct {
	sum      int
	mediator Mediator
}

func (a *account) add(amount int) {
	if amount > 0 {
		a.sum += amount
	} else {
		fmt.Println("incorrect amount")
	}
}

func (a *account) withdraw(amount int) {
	if amount > 0 || a.sum-amount < 0 {
		a.sum -= amount
	} else {
		fmt.Println("incorrect amount")
	}
}

// Send notifies a mediator for do operation (implements Account interface)
func (a *account) Send(id rune, amount int) {
	if a.mediator == nil {
		log.Println("mediator required")
		return
	}
	a.mediator.Notify('A', amount)
}

// GetSum gets an amount of money of the account (implements Account interface)
func (a *account) GetSum() int {
	return a.sum
}

// SetMediator sets a mediator for doing operation (implements Account interface)
func (a *account) SetMediator(mediator Mediator) {
	a.mediator = mediator
}

// Show shows all of fields of the account (implements Account interface)
func (a *account) Show() {
	fmt.Println(a)
}

//NewAccount ...
func NewAccount(amount int) Account {
	return &account{sum: amount}
}
