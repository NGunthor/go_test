package valid_parentheses

type ParenthesesString interface {
	IsValid() bool
}

type parenthesesString struct {
	s string
}

func (s *parenthesesString) IsValid() bool {
	stack := NewStack()
	for _, char := range s.s {
		if char == '(' || char == '[' || char == '{' {
			stack.Push(char)
		} else {
			switch char {
			case ')':
				if stack.Peak() != '(' {
					return false
				}
				stack.Pop()
			case '}':
				if stack.Peak() != '{' {
					return false
				}
				stack.Pop()
			case ']':
				if stack.Peak() != '[' {
					return false
				}
				stack.Pop()
			default:
				return false
			}
		}
	}
	if stack.Len() > 0 {
		return false
	}
	return true
}

func NewStr(s string) ParenthesesString {
	return &parenthesesString{s:s}
}
