func isValid(s string) bool {
	stack := make([]rune, 0)
	compare := map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	var t rune
	for _, i := range s {
		if i == '(' || i == '[' || i == '{' {
			stack = append(stack, i)
		} else {
			if len(stack) == 0 {
				return false
			}
			t, stack = stack[len(stack)-1], stack[:len(stack)-1]
			if compare[t] != i {
				return false
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true

}
