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
