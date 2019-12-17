package memento

import (
	"errors"
	"fmt"
)

type CareTaker interface {
	GetMem() (Memento, error)
	PushMem(mem Memento)
	Show()
}

type careTaker struct {
	memes []Memento
	size  int
}

func (c *careTaker) GetMem() (Memento, error) {
	var out Memento = nil
	var err error = nil
	if c.size > 0 {
		out = c.memes[0]
		c.memes = c.memes[1:]
		c.size--
	} else {
		err = errors.New("empty")
	}
	return out, err
}

func (c *careTaker) PushMem(mem Memento) {
	c.memes = append(c.memes, mem)
	c.size++
}

func (c *careTaker) Show() {
	for _, mem := range c.memes {
		fmt.Println(mem.getState())
	}
}

func NewCareTaker() CareTaker {
	return &careTaker{
		memes: make([]Memento, 0),
		size:  0,
	}
}
