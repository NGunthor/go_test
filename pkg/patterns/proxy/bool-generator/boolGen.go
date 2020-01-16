package bool_generator

import (
	"math/rand"
)

type BoolGenerator interface {
	Generate() bool
}

type boolGenerator struct {
	src       rand.Source
	cache     int64
	remaining int
}

func (b *boolGenerator) Generate() bool {
	if b.remaining == 0 {
		b.cache, b.remaining = b.src.Int63(), 63
	}

	result := b.cache&0x01 == 1
	b.cache >>= 1
	b.remaining--

	return result
}

func NewBoolGenerator(source rand.Source) BoolGenerator {
	return &boolGenerator{src: source}
}
