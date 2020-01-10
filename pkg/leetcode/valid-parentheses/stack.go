package valid_parentheses

// Stack provides functions for stack of runes
type Stack interface {
	Push(r rune)
	Pop() rune
	Len() int
	Peak() rune
}

type stack struct {
	values []rune
	size   int
}

// Push adds an element to stack
func (s *stack) Push(r rune) {
	s.values = append(s.values, r)
	s.size++
}

// Pop takes stack's element away
func (s *stack) Pop() rune {
	var out rune = 0
	if s.size > 0 {
		out = s.values[s.size-1]
		s.values = s.values[:s.size-1]
		s.size--
	}
	return out
}

// Len returns length of stack
func (s *stack) Len() int {
	return s.size
}

// Peak gets the highest stack's element
func (s *stack) Peak() rune {
	var out rune = 0
	if s.size > 0 {
		out = s.values[s.size-1]
	}
	return out
}

// NewStack ...
func NewStack() *stack {
	return &stack{[]rune{}, 0}
}
