package set

import "sort"

type set struct {
	items map[int]bool
	len   int
}

func (s *set) Add(elements ...int) *set {
	for _, element := range elements {
		s.items[element] = true
	}
	s.Len()
	return s
}

func (s *set) Len() int {
	s.len = len(s.items)
	return s.len
}

func (s *set) Delete(elements ...int) *set {
	for _, element := range elements {
		delete(s.items, element)
	}
	s.Len()
	return s
}

func (s set) ToSlice() []int {
	slice := make([]int, len(s.items))
	i := 0
	for k := range s.items {
		slice[i] = k
		i++
	}
	sort.Ints(slice)
	return slice
}

func (s set) Copy() set {
	newSet := NewSet()
	for k := range s.items {
		newSet.items[k] = true
	}
	newSet.Len()
	return *newSet
}

func NewSet(elements ...int) *set {
	out := set{items: make(map[int]bool, 0)}
	for _, element := range elements {
		out.Add(element)
	}
	out.Len()
	return &out
}

func (s set) Union(sets ...set) set {
	outSet := s.Copy()
	for _, set := range sets {
		for k, v := range set.items {
			outSet.items[k] = v
		}
	}
	outSet.Len()
	return outSet
}

func (s set) Difference(b set) set {
	outSet := s.Copy()
	for k := range outSet.items {
		if b.items[k] == true {
			outSet.Delete(k)
		}
	}
	outSet.Len()
	return outSet
}

func (s set) Intersection(sets ...set) set {
	outSet := s.Copy()
	var curSet set
	for _, set := range sets {
		curSet = outSet.Copy()
		outSet = *NewSet()
		for k := range set.items {
			if curSet.items[k] == true && set.items[k] == true {
				outSet.Add(k)
			}
		}
	}
	outSet.Len()
	return outSet
}

func (s set) IsSubset(subsets ...set) bool {
	for _, subset := range subsets {
		for k, v := range subset.items {
			if s.items[k] != v {
				return false
			}
		}
	}
	return true
}
