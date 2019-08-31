import "strconv"

func calPoints(ops []string) int {
	sums := 0
	stack := make([]int, 0)
	for _, i := range ops {
		if i == "C" {
			invalid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			sums = sums - invalid
		} else if i == "D" {
			valid := stack[len(stack)-1]
			stack = append(stack, valid*2)
			sums = sums + stack[len(stack)-1]
		} else if i == "+" {
			add := stack[len(stack)-1] + stack[len(stack)-2]
			stack = append(stack, add)
			sums = sums + add
		} else {
			d, _ := strconv.Atoi(i)
			stack = append(stack, d)
			sums = sums + stack[len(stack)-1]
		}

	}
	return sums
}
