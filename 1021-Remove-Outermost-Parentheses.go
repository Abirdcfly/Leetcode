func removeOuterParentheses(S string) string {
	opend := 0
	stack := make([]rune, 0)
	for _, i := range S {
		if i == '(' {
			opend++
			if opend > 1 {
				stack = append(stack, i)
			}

		} else {
			opend--
			if opend > 0 {
				stack = append(stack, i)
			}
		}
	}
	return string(stack)
}
