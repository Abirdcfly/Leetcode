func countBits(num int) []int {
	res := make([]int, num+1)
	for i := 1; i <= num; i++ {
		//res[i] = res[i>>1] + i&1 //也可以
		res[i] = res[i&(i-1)] + 1
	}
	return res
}

import "math/bits"
func countBits(num int) []int {
    res := make([]int,num+1)
    for i:=1;i<=num;i++{
        //res[i] = res[i>>1] + i&1
        //res[i] = res[i&(i-1)] + 1
        res[i] = bits.OnesCount64(uint64(i))
    }
    return res
}
