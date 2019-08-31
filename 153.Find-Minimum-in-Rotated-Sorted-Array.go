func findMin(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}
	if nums[len(nums)-1] > nums[0] {
		return nums[0]
	}
	s, e := 0, len(nums)-1
	m := (s + e) / 2
	for ; s <= e; m = (s + e) / 2 {
		if m+1 <= len(nums)-1 && nums[m] > nums[m+1] {
			return nums[m+1]
		}
		if m-1 >= 0 && nums[m-1] > nums[m] {
			return nums[m]
		}
		if nums[m] < nums[0] {
			e = m - 1
		} else {
			s = m + 1
		}
	}
	return -1
}
