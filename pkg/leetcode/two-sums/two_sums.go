package two_sums

// TwoSum the main algorithm function
func TwoSum(nums []int, target int) []int {
	theMap := map[int]int{}
	for i, num := range nums {
		theMap[num] = i
	}

	for i, num := range nums {
		if val, ok := theMap[target-num]; ok {
			if val <= i {
				continue
			}
			return []int{i, val}
		}
	}
	return nil
}
