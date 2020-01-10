package mac_factory

import "fmt"

type Button interface {
	Paint()
}

type macButton struct {
}

func (m *macButton) Paint() {
	fmt.Println("Mac Button painted")
}
