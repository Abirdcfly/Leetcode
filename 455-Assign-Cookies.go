import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	res := 0
	j := 0
	for _, i := range s {
		for k := j; k < len(g); k++ {
			if i >= g[k] {
				res++
				j = k + 1
				break
			}
		}
	}
	return res
}

func findContentChildren(g []int, s []int) int {
	if len(g) == 0 || len(s) == 0 {
		return 0
	}
	sort.Ints(g)
	sort.Ints(s)
	j := 0
	res := 0
	for i := range s {
		if j >= len(g) {
			break
		}
		if s[i] >= g[j] {
			res++
			j++
			continue
		}
	}
	return res
}
