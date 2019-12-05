package main

import "fmt"

func main() {

	m := map[int]int{1:0, 2:10}
	fmt.Println(m, m[0], m[1], m[1])
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sm := make(map[uint8]uint8, 0)
	tm := make(map[uint8]uint8, 0)
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
