package transports

import "fmt"

type Transport interface {
	Drive()
}

type auto struct {
}

// Drive performs drive work for Transport
func (a *auto) Drive() {
	fmt.Println("auto is driving")
}

func NewAuto() *auto {
	return &auto{}
}
