package transports

import "fmt"

// Animal provides Animal functional
type Animal interface {
	Move()
}

type horse struct {
}

// Move performs move work for Animal
func (h *horse) Move() {
	fmt.Println("horse is moving")
}

// NewHorse ...
func NewHorse() Animal {
	return &horse{}
}
