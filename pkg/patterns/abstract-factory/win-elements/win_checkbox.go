package win_elements

import "fmt"

type winCheckbox struct {
}

// Paint paints windows checkbox
func (w *winCheckbox) Paint() {
	fmt.Println("Windows Checkbox painted")
}

func NewWinCheckbox() Painter {
	return &winCheckbox{}
}
