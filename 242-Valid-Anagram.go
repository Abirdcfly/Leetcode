import (
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sm := make(map[rune]int)
	for _, i := range s {
		sm[i] += 1
	}
	for _, i := range t {
		if v, exist := sm[i]; !exist {
			return false
		} else {
			sm[i] = v - 1
		}
	}
	for _, v := range sm {
		if v != 0 {
			return false
		}
	}
	return true
}

// 以上的构建一个map 第二个来递减第一个可以，也可以构建2个map来比较，或者排序来处理。
func isAnagram(s string, t string) bool {
	ss := sortstring(s)
	st := sortstring(t)
	return ss == st
}

func sortstring(s string) string {
	news := strings.Split(s, "")
	sort.Strings(news)
	return strings.Join(news, "")
}

