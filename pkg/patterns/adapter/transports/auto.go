package transports

import "fmt"

// Transport provides transports functional
type Transport interface {
	Drive()
}

type auto struct {
}

// Drive performs drive work for Transport
func (a *auto) Drive() {
	fmt.Println("auto is driving")
}

// NewAuto ...
func NewAuto() *auto {
	return &auto{}
}