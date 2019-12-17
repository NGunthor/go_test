package mediator

import (
	"fmt"
	"log"
)

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

func (a *account) Send(id rune, amount int) {
	if a.mediator == nil {
		log.Println("mediator required")
		return
	}
	a.mediator.Notify('A', amount)
}

func (a *account) GetSum() int {
	return a.sum
}

func (a *account) SetMediator(mediator Mediator) {
	a.mediator = mediator
}

func (a *account) Show() {
	fmt.Println(a)
}

func NewAccount(amount int) Account {
	return &account{sum: amount}
}
