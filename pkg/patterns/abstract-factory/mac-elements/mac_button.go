package mac_elements

import "fmt"

type macButton struct {
}

// Paint paints mac button (implements Painter interface)
func (m *macButton) Paint() {
	fmt.Println("Mac Button painted")
}

// NewMacButton ...
func NewMacButton() Painter {
	return &macButton{}
}
