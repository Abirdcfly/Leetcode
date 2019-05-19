import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	f := func(s string) string {
		ns := strings.Split(s, "")
		sort.Strings(ns)
		return strings.Join(ns, "")
	}
	res := make(map[string][]string)
	for _, i := range strs {
		sorti := f(i)
		if v, exist := res[sorti]; exist {
			v = append(v, i)
			res[sorti] = v
		} else {
			res[sorti] = []string{i}
		}
	}
	r := make([][]string, 0)
	for _, v := range res {
		r = append(r, v)
	}
	return r
}
