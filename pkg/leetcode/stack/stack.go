package stack

// Stack provides functions for stack of ints
type Stack interface {
	Push(r int)
	Pop() int
	Len() int
	Peak() int
}

type stack struct {
	values []int
	size   int
}

// Push adds an element to stack
func (s *stack) Push(r int) {
	s.values = append(s.values, r)
	s.size++
}

// Pop takes stack's element away
func (s *stack) Pop() int {
	var out int = 0
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
func (s *stack) Peak() int {
	var out int = 0
	if s.size > 0 {
		out = s.values[s.size-1]
	}
	return out
}

// NewStack ...
func NewStack() *stack {
	return &stack{[]int{}, 0}
}
