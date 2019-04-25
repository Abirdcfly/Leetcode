func backspaceCompare(S string, T string) bool {
	getres := func(s string) string {
		stack := make([]rune, 0)
		for _, i := range s {
			if i == '#' {
				if len(stack) != 0 {
					stack = stack[:len(stack)-1]
				}
			} else {
				stack = append(stack, i)
			}
		}
		return string(stack)
	}
	return getres(S) == getres(T)
}
