package win_factory

import "fmt"

type checkbox interface {
	Paint()
}

type winCheckbox struct {
}

func (w *winCheckbox) Paint() {
	fmt.Println("Windows Checkbox painted")
}
