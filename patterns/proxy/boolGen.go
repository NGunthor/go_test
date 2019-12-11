package proxy

import (
	"math/rand"
	"time"
)

type boolgen struct {
	src       rand.Source
	cache     int64
	remaining int
}

func (b *boolgen) Bool() bool {
	if b.remaining == 0 {
		b.cache, b.remaining = b.src.Int63(), 63
	}

	result := b.cache&0x01 == 1
	b.cache >>= 1
	b.remaining--

	return result
}

func NewBoolGen() *boolgen {
	return &boolgen{src: rand.NewSource(time.Now().UnixNano())}
}
