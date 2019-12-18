package stack

type stack struct {
	values []int
	size   int
}

func (s *stack) Push(r int) {
	s.values = append(s.values, r)
	s.size++
}

func (s *stack) Pop() int {
	var out int = 0
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

func (s *stack) Peak() int {
	var out int = 0
	if s.size > 0 {
		out = s.values[s.size-1]
	}
	return out
}

func NewStack() *stack {
	return &stack{[]int{}, 0}
}
