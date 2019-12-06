package set

type addTests struct {
	values []int
	result []int
}

type deleteTests struct {
	start  set
	values []int
	result []int
}

type unionTests struct {
	sets   []set
	result []int
}

type intersectionTests struct {
	start  set
	sets   []set
	result []int
}

type differenceTests struct {
	a      set
	b      set
	result []int
}

type subsetTests struct {
	set    set
	subset set
	result bool
}
