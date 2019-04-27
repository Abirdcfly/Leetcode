func findrotated(nums []int) int {
	//153. Find Minimum in Rotated Sorted Array
	if len(nums) < 2 {
		return 0
	}
	s, e := 0, len(nums)-1
	//1 < 2 < 3 < 4 < 5 < 7. Already sorted array.
	if nums[e] > nums[0] {
		return 0
	}
	m := (s + e) / 2
	for ; s <= e; m = (s + e) / 2 {
		if nums[m] > nums[m+1] {
			return m + 1
		}
		if nums[m-1] > nums[m] {
			return m
		}
		if nums[m] > nums[0] {
			s = m + 1
		} else {
			e = m - 1
		}
	}
	return -1
}

func search(nums []int, target int) int {
	m := findrotated(nums)
	s1, e1 := 0, len(nums)-1
	m1 := (s1 + e1) / 2
	for ; s1 <= e1; m1 = (s1 + e1) / 2 {
		realm := (m1 + m) % len(nums)
		if nums[realm] > target {
			e1 = m1 - 1
		} else if nums[realm] < target {
			s1 = m1 + 1
		} else {
			return realm
		}
	}
	return -1
}
