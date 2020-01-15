package two_sums

type Numbers interface {
	TwoSum() []int
}

type numbers struct {
	nums   []int
	target int
}

func (n *numbers) TwoSum() []int {
	theMap := map[int]int{}
	for i, num := range n.nums {
		theMap[num] = i
	}

	for i, num := range n.nums {
		if val, ok := theMap[n.target-num]; ok {
			if val <= i {
				continue
			}
			return []int{i, val}
		}
	}
	return nil
}

func NewNumbers(nums []int, target int) Numbers {
	return &numbers{
		nums:   nums,
		target: target,
	}
}
