import "sort"

func findLengthOfLCIS(nums []int) int {
	l := len(nums)
	if l == 0 || l == 1 {
		return l
	}
	m := make([]int, len(nums))
	//m[i] 必须包含i时的增长序列长度
	m[0] = 1
	for i := 1; i < l; i++ {
		if nums[i] > nums[i-1] {
			m[i] = m[i-1] + 1
		} else {
			m[i] = 1
		}
	}
	sort.Ints(m)
	return m[l-1]
}

func findLengthOfLCIS(nums []int) int {
	l := len(nums)
	if l == 0 || l == 1 {
		return l
	}
	ml, m, res := 1, 0, 0 //必须包含前一个数字时最大长度，必须包含当前数字时当前最大长度，结果
	for i := 1; i < l; i++ {
		if nums[i] > nums[i-1] {
			m = ml + 1
		} else {
			m = 1
		}
		if m > res {
			res = m
		}
		ml = m
	}
	return res
}


