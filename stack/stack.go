package stack

type stack struct {
	values []rune
	size   int
}

func (s *stack) Push(r rune) {
	s.values = append(s.values, r)
	s.size++
}

func (s *stack) Pop() rune {
	var out rune = 0
	if s.size > 0 {
		out = s.values[s.size-1]
		s.values = s.values[:s.size-1]
		s.size--
	}
	return out
}

func (s *stack) Len() int {
	return s.size
}

func (s *stack) Peak() rune {
	var out rune = 0
	if s.size > 0 {
		out = s.values[s.size-1]
	}
	return out
}

func NewStack() *stack {
	return &stack{[]rune{}, 0}
}
