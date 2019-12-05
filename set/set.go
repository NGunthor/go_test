package set

type set struct {
	Items map[int]bool
	len int
}

func Add(s *set, elements ...int) {
	for _, element := range elements {
		s.Items[element] = true
	}
	s.len = len(s.Items)
}

func Delete(s *set, elements ...int) {
	for _, element := range elements {
		delete(s.Items, element)
	}
	s.len = len(s.Items)
}

func Copy(set *set) set {
	newSet := Make()
	for k := range set.Items {
		newSet.Items[k] = true
	}
	newSet.len = len(newSet.Items)
	return newSet
}

func Make() set {
	return set{map[int]bool{}, 0}
}

func Union(sets ...*set) set {
	outSet := Make()
	for _, set := range sets {
		for k, v := range set.Items {
			outSet.Items[k] = v
		}
	}
	outSet.len = len(outSet.Items)
	return outSet
}

func Difference(a *set, b *set)  set{
	outSet := Copy(a)
	for k := range outSet.Items {
		if b.Items[k] == true {
			Delete(&outSet, k)
		}
	}
	outSet.len = len(outSet.Items)
	return outSet
}

func Intersection(sets ...*set) set {
	outSet := Make()
	currentSet := Copy(sets[0])
	for i, set := range sets {
		for k, v := range set.Items {
			if currentSet.Items[k] == true && v == true && i != 0 {
				Add(&outSet, k)
			}
		}
	}
	outSet.len = len(outSet.Items)
	return outSet
}

func IsSubset(source *set, subsets ...*set) bool{
	for _, subset := range subsets {
		for k, v := range subset.Items {
			if source.Items[k] != v {
				return false
			}
		}
	}
	return true
}
