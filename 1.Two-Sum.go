func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, num := range nums {
		if index2, exist := m[target-num]; exist {
			return []int{index, index2}
		}
		m[num] = index
	}
	return nil
}
