func findKthLargest(nums []int, k int) int {
	return QuickSort(nums, 0, len(nums)-1, k)
}

func QuickSort(a []int, s, e, K int) int {
	if s == e {
		return a[s]
	}
	i := partition(a, s, e)
	//s,e,i指的都是下标。K是第几个数。中间差1
	if i+1 > K {
		return QuickSort(a, s, i-1, K)
	} else if i+1 < K {
		return QuickSort(a, i+1, e, K)
	} else {
		return a[i]
	}
}

func partition(a []int, s, e int) int {
	pivot := a[e]
	i, j := s, s
	for ; j <= e; j++ {
		if a[j] > pivot {
			a[j], a[i] = a[i], a[j]
			i++
		}
	}
	a[i], a[e] = a[e], a[i]
	return i
}
