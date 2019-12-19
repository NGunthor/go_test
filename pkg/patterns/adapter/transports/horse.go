package transports

import "fmt"

type Animal interface {
	Move()
}

type horse struct {
}

// Move performs move work for Animal
func (h *horse) Move() {
	fmt.Println("horse is moving")
}

func NewHorse() *horse {
	return &horse{}
}
