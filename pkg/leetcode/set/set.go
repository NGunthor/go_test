package set

import "sort"

// Set provides all main methods for set struct
type Set interface {
	Add(elements ...int)
	Len() int
	Delete(elements ...int)
	ToSlice() []int
	Copy() set
	Union(sets ...set) set
	Difference(b set) set
	Intersection(sets ...set) set
	IsSubset(subsets ...set) bool
}

type set struct {
	items map[int]bool
	len   int
}

// Add adds elements to actual set object
func (s *set) Add(elements ...int) {
	for _, element := range elements {
		s.items[element] = true
	}
	s.Len()
}

// Len return size for actual set object
func (s *set) Len() int {
	s.len = len(s.items)
	return s.len
}

// Delete deletes elements from actual set object
func (s *set) Delete(elements ...int) {
	for _, element := range elements {
		delete(s.items, element)
	}
	s.Len()
}

// ToSlice converts set's values to slice and returns the slice
func (s *set) ToSlice() []int {
	slice := make([]int, len(s.items))
	i := 0
	for k := range s.items {
		slice[i] = k
		i++
	}
	sort.Ints(slice)
	return slice
}

// Copy returns copy for actual set object
func (s *set) Copy() set {
	newSet := NewSet()
	for k := range s.items {
		newSet.items[k] = true
	}
	newSet.Len()
	return *newSet
}

// Union returns set that is a union of all passed sets
func (s *set) Union(sets ...set) set {
	outSet := s.Copy()
	for _, set := range sets {
		for k, v := range set.items {
			outSet.items[k] = v
		}
	}
	outSet.Len()
	return outSet
}

// Difference returns set that is a difference of all passed sets
func (s *set) Difference(b set) set {
	outSet := s.Copy()
	for k := range outSet.items {
		if b.items[k] == true {
			outSet.Delete(k)
		}
	}
	outSet.Len()
	return outSet
}

// Intersection returns set that is an intersection of all passed sets
func (s *set) Intersection(sets ...set) set {
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

// IsSubset returns true if all passed sets are subsets for actual set
func (s *set) IsSubset(subsets ...set) bool {
	for _, subset := range subsets {
		for k, v := range subset.items {
			if s.items[k] != v {
				return false
			}
		}
	}
	return true
}

// NewSet ...
func NewSet(elements ...int) *set {
	out := set{items: make(map[int]bool, 0)}
	for _, element := range elements {
		out.Add(element)
	}
	out.Len()
	return &out
}
