package win_factory

import "fmt"

type button interface {
	Paint()
}

type winButton struct {
}

func (w *winButton) Paint() {
	fmt.Println("Windows Button painted")
}
