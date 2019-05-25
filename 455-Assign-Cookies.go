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
