package mac_elements

import "fmt"

type macCheckbox struct {
}

// Paint paints mac checkbox
func (m *macCheckbox) Paint() {
	fmt.Println("Mac Checkbox painted")
}

// NewMacCheckbox ...
func NewMacCheckbox() Painter {
	return &macCheckbox{}
}
