package is_isomorphic

// Strings provides interface for strings struct
type Strings interface {
	IsIsomorphic() bool
}

type strings struct {
	string1 string
	string2 string
}

// IsIsomorphic is a main algorithm function
func (s *strings) IsIsomorphic() bool {
	if len(s.string1) != len(s.string2) {
		return false
	}
	sm := make(map[uint8]uint8, len(s.string1))
	tm := make(map[uint8]uint8, len(s.string2))
	for i := range s.string1 {
		sm[s.string1[i]] = s.string2[i]
		tm[s.string2[i]] = s.string1[i]
	}
	for i := range s.string1 {
		if s.string2[i] != sm[s.string1[i]] || s.string1[i] != tm[s.string2[i]] {
			return false
		}
	}
	return true
}

// NewStrings ...
func NewStrings(str1, str2 string) Strings {
	return &strings{
		string1: str1,
		string2: str2,
	}
}
