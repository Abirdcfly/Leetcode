import "sort"

func lengthOfLIS(nums []int) int {
	l := len(nums)
	if l == 0 || l == 1 {
		return l
	}
	m := make([]int, len(nums))
	//m 包含i时最长递增子序列
	m[0] = 1
	for i := 1; i < l; i++ {
		t, t1 := 0, 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				t1 = m[j]
				if t1 > t {
					t = t1
				}
			} //找到0->i-1中 小于nums[i] 且最大的长度
			m[i] = t + 1
		}
	}
	sort.Ints(m)
	return m[l-1]
}
