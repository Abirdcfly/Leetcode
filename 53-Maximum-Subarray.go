import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	f := make([]int, len(nums)) // subarray max val that must end with nums[i]
	f[0] = nums[0]
	max := float64(f[0])
	for i := 1; i < len(nums); i++ {
		if f[i-1] > 0 {
			f[i] = f[i-1] + nums[i]
		} else {
			f[i] = nums[i]
		}
		max = math.Max(max, float64(f[i]))
	}
	fmt.Println(f)
	return int(max)
}
