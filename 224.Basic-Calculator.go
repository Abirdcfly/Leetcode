import "strconv"

//https://www.youtube.com/watch?v=9c0WHgIsk5g
type Flag struct {
	f string
	r int
}

func calculate(s string) int {
	nums := make([]int, 0)
	flags := make([]Flag, 0)
	rank := 0
	flag := Flag{}
	num1, num2, num := 0, 0, 0
	getelement := func(s string) (res []string) {
		ele := make([]rune, 0)
		for _, ii := range s {
			if ii == ' ' {
				continue
			}
			if ii == '(' || ii == ')' || ii == '+' || ii == '-' {
				if len(ele) > 0 {
					res = append(res, string(ele))
					ele = ele[:0]
				}
				res = append(res, string(ii))
			} else {
				ele = append(ele, ii)
			}
		}
		if len(ele) > 0 {
			res = append(res, string(ele))
		}
		return res
	}
	for _, i := range getelement(s) {
		if i == "+" || i == "-" {
			for len(flags) > 0 && flags[len(flags)-1].r >= rank {
				flag, flags = flags[len(flags)-1], flags[:len(flags)-1]
				num2, nums = nums[len(nums)-1], nums[:len(nums)-1]
				num1, nums = nums[len(nums)-1], nums[:len(nums)-1]
				if flag.f == "+" {
					num = num1 + num2
					nums = append(nums, num)
				} else {
					num = num1 - num2
					nums = append(nums, num)
				}
			}
			flags = append(flags, Flag{f: i, r: rank})
		} else if i == "(" {
			rank++
		} else if i == ")" {
			rank--
		} else {
			n, _ := strconv.Atoi(i)
			nums = append(nums, n)
		}
	}
	for len(flags) > 0 {
		flag, flags = flags[len(flags)-1], flags[:len(flags)-1]
		num2, nums = nums[len(nums)-1], nums[:len(nums)-1]
		num1, nums = nums[len(nums)-1], nums[:len(nums)-1]
		if flag.f == "+" {
			num = num1 + num2
			nums = append(nums, num)
		} else {
			num = num1 - num2
			nums = append(nums, num)
		}
	}
	return nums[0]
}
