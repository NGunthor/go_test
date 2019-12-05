package isIsomorphic

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sm := make(map[uint8]uint8, len(s))
	tm := make(map[uint8]uint8, len(t))
	for i := range s {
		sm[s[i]] = t[i]
		tm[t[i]] = s[i]
	}
	for i := range s {
		if t[i] != sm[s[i]] ||  s[i] != tm[t[i]]{
			return false
		}
	}
	return true
}
