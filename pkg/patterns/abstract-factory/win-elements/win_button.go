package win_elements

import "fmt"

type winButton struct {
}

// Paint paints windows button
func (w *winButton) Paint() {
	fmt.Println("Windows Button painted")
}

// NewWinButton ...
func NewWinButton() Painter {
	return &winButton{}
}
