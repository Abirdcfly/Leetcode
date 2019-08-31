
//暴力求解 O((n-k+1)*k)
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nums
	}
	res := make([]int, 0)
	for i := 0; i+k <= len(nums); i++ {
		res = append(res, max(nums, i, i+k))
	}
	return res
}

func max(nums []int, b, e int) int {
	r := nums[b]
	for i := b; i < e; i++ {
		if nums[i] > r {
			r = nums[i]
		}
	}
	return r
}

//所谓单调队列
func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0)
	window := make([]int, 0)
	if len(nums) == 0 {
		return res
	}
	for index, i := range nums {
		if index >= k && window[0] <= index-k {
			window = window[1:]
		}
		for len(window) > 0 && nums[window[len(window)-1]] <= i {
			window = window[:len(window)-1]
		}
		window = append(window, index)
		if index >= k-1 {
			res = append(res, nums[window[0]])
		}
	}
	return res
}
