import "strconv"

func evalRPN(tokens []string) int {
	nums := make([]int, 0)
	x1 := 0
	num1, num2, newnum := 0, 0, 0
	for i := 0; i < len(tokens); i++ {
		x := tokens[i]
		if x == "+" || x == "-" || x == "*" || x == "/" {
			num1, nums = nums[len(nums)-1], nums[:len(nums)-1]
			num2, nums = nums[len(nums)-1], nums[:len(nums)-1]
			if x == "+" {
				newnum = num2 + num1
			} else if x == "-" {
				newnum = num2 - num1
			} else if x == "*" {
				newnum = num2 * num1
			} else if x == "/" {
				newnum = num2 / num1
			}
			nums = append(nums, newnum)
		} else {
			x1, _ = strconv.Atoi(x)
			nums = append(nums, x1)
		}
	}
	return nums[0]
}
