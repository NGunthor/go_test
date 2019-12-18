package valid_parentheses

func IsValid(s string) bool {
	stack := NewStack()
	for _, char := range s {
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
