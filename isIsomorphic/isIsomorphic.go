package main

import (
	"fmt"
)

func main() {

	s := "egg"
	t := "add"

	fmt.Println(isIsomorphic(s, t))
}

func isIsomorphic(s string, t string) bool {
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
