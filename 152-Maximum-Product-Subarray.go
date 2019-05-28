import "sort"

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	m := make([][][]int, 2)
	for i := range m {
		m[i] = make([][]int, 2)
		for j := range m[i] {
			m[i][j] = make([]int, 2)
		}
	}
	//m[i][j][k] i:ith,last 本次和上一次 j 0负max 1max k0当前在里面 1当前不在里面
	if t := nums[0]; t > 0 {
		m[0][1][0] = t
	} else if t < 0 {
		m[0][0][0] = t
	}

	for i := 1; i < len(nums); i++ {
		x := nums[i]
		ith := i % 2
		last := (i - 1 + 2) % 2
		if x > 0 {
			m[ith][0][0] = m[last][0][0] * x
			m[ith][1][0] = max(x, m[last][1][0]*x)
		} else if x < 0 {
			m[ith][0][0] = min(x, m[last][1][0]*x)
			m[ith][1][0] = max(0, m[last][0][0]*x)
		} else {
			m[ith][0][0] = 0
			m[ith][1][0] = 0
		}
		m[ith][0][1] = min(m[last][0][1], m[last][0][0])
		m[ith][1][1] = max(m[last][1][1], m[last][1][0])
	}
	return max(m[(len(nums)-1+2)%2][1][0], m[(len(nums)-1+2)%2][1][1])
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	m := make([]int, len(nums))
	nm := make([]int, len(nums))
	//m[i] i存在子序列里时，最大值
	//nm[i] i存在子序列里时，负最大值
	if nums[0] >= 0 {
		m[0] = nums[0]
	} else {
		nm[0] = nums[0]
	}
	for i := 1; i < len(nums); i++ {
		x := nums[i]
		if x >= 0 {
			m[i] = max(m[i-1]*x, x)
			nm[i] = nm[i-1] * x
		} else {
			m[i] = nm[i-1] * x
			nm[i] = min(m[i-1]*x, x)
		}
	}
	sort.Ints(m)
	return m[len(nums)-1]
}

func max(a ...int) int {
	res := a[0]
	for i := range a {
		if a[i] > res {
			res = a[i]
		}
	}
	return res
}

func min(a ...int) int {
	res := a[0]
	for i := range a {
		if a[i] < res {
			res = a[i]
		}
	}
	return res
}

func maxProduct(nums []int) int {
	res := 0
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	m := make([]int, 2)
	nm := make([]int, 2)
	//m[i] i存在子序列里时，最大值
	//nm[i] i存在子序列里时，负最大值
	m[0], nm[0], res = nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		x := nums[i]
		ith, last := i%2, (i+1)%2
		m[ith] = max(m[last]*x, nm[last]*x, x)
		nm[ith] = min(m[last]*x, nm[last]*x, x)
		res = max(m[ith], res)
	}
	return res
}
