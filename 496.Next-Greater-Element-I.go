func nextGreaterElement(nums1 []int, nums2 []int) []int {
	greater := make(map[int]int)
	s := make([]int, 0)
	num := 0
	for _, i := range nums2 {
		for len(s) > 0 && i > s[len(s)-1] {
			num, s = s[len(s)-1], s[:len(s)-1]
			greater[num] = i
		}
		s = append(s, i)
	}
	res := make([]int, 0)
	for _, i := range nums1 {
		if v, ok := greater[i]; ok {
			res = append(res, v)
		} else {
			res = append(res, -1)
		}
	}
	return res
}
