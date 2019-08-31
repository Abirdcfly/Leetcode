import (
	"fmt"
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

//方法都是殊途同归，核心是怎么找到同字母单词的唯一标记，可以是排序后的结果，也可以是别的，排序的话时间复杂度比算[0] * 26 大
func groupAnagrams(strs []string) [][]string {
	f := func(s string) string {
		res := make([]int, 26)
		for _, r := range s {
			res[int(r-'a')]++
		}
		return fmt.Sprint(res)
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
