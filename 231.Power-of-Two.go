import "math"

func isPowerOfTwo(n int) bool {
	if n < 0 {
		return false
	}
	i := 0
	for {
		if test := int(math.Pow(2.0, float64(i))); test == n {
			return true
		} else if test > n {
			break
		}
		i++
	}
	return false
}

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return n&(n-1) == 0
}

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	for n%2 != 1 {
		n = n / 2
	}
	return n == 1
}

